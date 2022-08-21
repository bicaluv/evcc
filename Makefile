.PHONY: default all clean install install-ui ui assets docs lint test-ui lint-ui test build test-release release
.PHONY: docker publish-testing publish-nightly publish-release
.PHONY: prepare-image image-rootfs image-update
.PHONY: soc

# build vars
TAG_NAME := $(shell test -d .git && git describe --abbrev=0 --tags)
SHA := $(shell test -d .git && git rev-parse --short HEAD)
COMMIT := $(SHA)
# hide commit for releases
ifeq ($(RELEASE),1)
    COMMIT :=
endif
VERSION := $(if $(TAG_NAME),$(TAG_NAME),$(SHA))
BUILD_DATE := $(shell date -u '+%Y-%m-%d_%H:%M:%S')
BUILD_TAGS := -tags=release
LD_FLAGS := -X github.com/evcc-io/evcc/server.Version=$(VERSION) -X github.com/evcc-io/evcc/server.Commit=$(COMMIT) -s -w
BUILD_ARGS := -ldflags='$(LD_FLAGS)'

# docker
DOCKER_IMAGE := andig/evcc
PLATFORM := linux/amd64,linux/arm64,linux/arm/v6

# gokrazy image
IMAGE_FILE := evcc_$(TAG_NAME).image
IMAGE_ROOTFS := evcc_$(TAG_NAME).rootfs
IMAGE_OPTIONS := -hostname evcc -http_port 8080 github.com/gokrazy/serial-busybox github.com/gokrazy/breakglass github.com/evcc-io/evcc

# deb
PACKAGES = ./release

default: build

all: clean install install-ui ui assets lint test-ui lint-ui test build

clean:
	rm -rf dist/

install:
	go install $$(go list -f '{{join .Imports " "}}' tools.go)

install-ui:
	npm ci

ui:
	npm run build

assets:
	go generate ./...

docs:
	go generate github.com/evcc-io/evcc/util/templates/...

lint:
	golangci-lint run

lint-ui:
	npm run lint

test-ui:
	npm test

test:
	@echo "Running testsuite"
	go test $(BUILD_TAGS) ./...

build:
	@echo Version: $(VERSION) $(SHA) $(BUILD_DATE)
	go build -v $(BUILD_TAGS) $(BUILD_ARGS)

release-test:
	goreleaser --snapshot --skip-publish --rm-dist

release:
	goreleaser --rm-dist

docker:
	@echo Version: $(VERSION) $(SHA) $(BUILD_DATE)
	docker build --tag $(DOCKER_IMAGE):testing .

publish-testing:
	@echo Version: $(VERSION) $(SHA) $(BUILD_DATE)
	docker buildx build --platform $(PLATFORM) --tag $(DOCKER_IMAGE):testing --push .

publish-nightly:
	@echo Version: $(VERSION) $(SHA) $(BUILD_DATE)
	docker buildx build --platform $(PLATFORM) --tag $(DOCKER_IMAGE):nightly .

publish-release:
	@echo Version: $(VERSION) $(SHA) $(BUILD_DATE)
	docker buildx build --build-arg RELEASE=1 --platform $(PLATFORM) --tag $(DOCKER_IMAGE):latest .

apt-nightly:
	$(foreach file, $(wildcard $(PACKAGES)/*.deb), \
		cloudsmith push deb evcc/unstable/any-distro/any-version $(file); \
	)

apt-release:
	$(foreach file, $(wildcard $(PACKAGES)/*.deb), \
		cloudsmith push deb evcc/stable/any-distro/any-version $(file); \
	)

# gokrazy image
prepare-image:
	go install github.com/gokrazy/tools/cmd/gokr-packer@latest
	mkdir -p flags/github.com/gokrazy/breakglass
	echo "-forward=private-network" > flags/github.com/gokrazy/breakglass/flags.txt
	mkdir -p buildflags/github.com/evcc-io/evcc
	echo "$(BUILD_TAGS),gokrazy" > buildflags/github.com/evcc-io/evcc/buildflags.txt
	echo "$(BUILD_ARGS)" >> buildflags/github.com/evcc-io/evcc/buildflags.txt

image:
	gokr-packer -overwrite=$(IMAGE_FILE) -target_storage_bytes=1258299392 $(IMAGE_OPTIONS)
	loop=$$(sudo losetup --find --show -P $(IMAGE_FILE)); sudo mkfs.ext4 $${loop}p4
	gzip -f $(IMAGE_FILE)

image-rootfs:
	gokr-packer -overwrite_root=$(IMAGE_ROOTFS) $(IMAGE_OPTIONS)
	gzip -f $(IMAGE_ROOTFS)

image-update:
	gokr-packer -update yes $(IMAGE_OPTIONS)

soc:
	@echo Version: $(VERSION) $(SHA) $(BUILD_DATE)
	go build $(BUILD_TAGS) $(BUILD_ARGS) github.com/evcc-io/evcc/cmd/soc
	
raspberrypi:
	@echo Version: $(VERSION) $(BUILD_DATE)
	# mm make all without tests
	make clean install install-ui ui assets build

	# make mac version
	mv evcc evcc_mac

	# make raspberry version
	env GOOS=linux GOARCH=arm go build -v $(BUILD_TAGS) $(BUILD_ARGS)
	
	# stop already running service and copy new evcc to raspberry
	ssh pi@raspberrypi 'sudo systemctl stop evcc.service; mv ~/bin/evcc ~/bin/evcc_prev'
	scp -B evcc pi@raspberrypi:~/bin/evcc
	# chmod and start service
	ssh pi@raspberrypi 'sudo sysctl -w net.ipv4.ping_group_range="0 2147483647"; sudo chmod 0755 ~/bin/evcc; sudo systemctl start evcc.service'

sync-with-andig-and-deploy:
	# show commits in browser
	open https://github.com/evcc-io/evcc/commits/master
	@echo "Proceed with build? (Y|n)"
	@read line; if [ "$$line" != "Y" ]; then echo aborting; exit 1 ; fi
	@echo building...

	# follwoing may show error if remote upstream alread configured
	# git remote add upstream https://github.com/andig/evcc.git
	# show github remote streams
	git remote -v
	# get new changes
	git fetch upstream
	# local branch
	git checkout master
	# get changes
	git merge upstream/master --no-edit --no-commit
	
	# build and deploy
	make raspberrypi
	
	# delete dist folder changes
	git restore dist
	git clean -d -f dist
	# push new files
	git add -A
	git commit -am "merge with evcc upstream"
	git push origin master
	# update sync detection file
	ssh pi@raspberrypi 'sudo echo -1 > ~/etc/check_fork.state'
