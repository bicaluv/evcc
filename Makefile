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
DOCKER_IMAGE := evcc/evcc
PLATFORM := linux/amd64,linux/arm64,linux/arm/v6

# gokrazy image
IMAGE_FILE := evcc_$(TAG_NAME).image
IMAGE_OPTIONS := -hostname evcc -http_port 8080 github.com/gokrazy/serial-busybox github.com/gokrazy/breakglass github.com/gokrazy/mkfs github.com/gokrazy/wifi github.com/evcc-io/evcc

# deb
PACKAGES = ./release

# asn1-patch
GOROOT := $(shell go env GOROOT)
CURRDIR := $(shell pwd)

default:: ui build

all:: clean install install-ui ui assets lint test-ui lint-ui test build

clean::
	rm -rf dist/

install::
	go install $$(go list -f '{{join .Imports " "}}' tools.go)

install-ui::
	npm ci

ui::
	npm run build

assets::
	go generate ./...

docs::
	go generate github.com/evcc-io/evcc/util/templates/...

lint::
	golangci-lint run

lint-ui::
	npm run lint

test-ui::
	npm test

test::
	@echo "Running testsuite"
	CGO_ENABLED=0 go test $(BUILD_TAGS) ./...

porcelain::
	gofmt -w -l $$(find . -name '*.go')
	go mod tidy
	test -z "$$(git status --porcelain)" || (git status; git diff; false)

build::
	@echo Version: $(VERSION) $(SHA) $(BUILD_DATE)
	CGO_ENABLED=0 go build -v $(BUILD_TAGS) $(BUILD_ARGS)

snapshot:
	goreleaser --snapshot --skip-publish --rm-dist

release::
	goreleaser --rm-dist

docker::
	@echo Version: $(VERSION) $(SHA) $(BUILD_DATE)
	docker buildx build --platform $(PLATFORM) --tag $(DOCKER_IMAGE):testing .

publish-testing::
	@echo Version: $(VERSION) $(SHA) $(BUILD_DATE)
	docker buildx build --platform $(PLATFORM) --tag $(DOCKER_IMAGE):testing --push .

publish-nightly::
	@echo Version: $(VERSION) $(SHA) $(BUILD_DATE)
	docker buildx build --platform $(PLATFORM) --tag $(DOCKER_IMAGE):nightly --push .

publish-release::
	@echo Version: $(VERSION) $(SHA) $(BUILD_DATE)
	docker buildx build --build-arg RELEASE=1 --platform $(PLATFORM) --tag $(DOCKER_IMAGE):latest --tag $(DOCKER_IMAGE):$(VERSION) --push .

apt-nightly::
	$(foreach file, $(wildcard $(PACKAGES)/*.deb), \
		cloudsmith push deb evcc/unstable/any-distro/any-version $(file); \
	)

apt-release::
	$(foreach file, $(wildcard $(PACKAGES)/*.deb), \
		cloudsmith push deb evcc/stable/any-distro/any-version $(file); \
	)

# gokrazy image
gokrazy::
	go install github.com/gokrazy/tools/cmd/gokr-packer@main
	mkdir -p flags/github.com/gokrazy/breakglass
	echo "-forward=private-network" > flags/github.com/gokrazy/breakglass/flags.txt
	mkdir -p flags/github.com/evcc-io/evcc
	echo "--sqlite=/perm/evcc.db" > flags/github.com/evcc-io/evcc/flags.txt
	mkdir -p env/github.com/evcc-io/evcc
	echo "EVCC_NETWORK_PORT=80" > env/github.com/evcc-io/evcc/env.txt
	mkdir -p buildflags/github.com/evcc-io/evcc
	echo "$(BUILD_TAGS),gokrazy" > buildflags/github.com/evcc-io/evcc/buildflags.txt
	echo "-ldflags=$(LD_FLAGS)" >> buildflags/github.com/evcc-io/evcc/buildflags.txt
	gokr-packer -hostname evcc -http_port 8080 -overwrite=$(IMAGE_FILE) -target_storage_bytes=1258299392 $(IMAGE_OPTIONS)
	# gzip -f $(IMAGE_FILE)

gokrazy-run::
	MACHINE=arm64 IMAGE_FILE=$(IMAGE_FILE) ./packaging/gokrazy/run.sh

gokrazy-update::
	gokr-packer -update yes $(IMAGE_OPTIONS)

soc::
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
	
	# copy new evcc to raspberry
	scp -B evcc pi@raspberrypi:~/bin/evcc_new

	# prepare restart of service
	ssh pi@raspberrypi 'sudo sysctl -w net.ipv4.ping_group_range="0 2147483647"; sudo chmod 0755 ~/bin/evcc_new;'

	# rename new/old evcc and restart already running service 
	ssh pi@raspberrypi 'mv ~/bin/evcc ~/bin/evcc_prev; mv ~/bin/evcc_new ~/bin/evcc; sudo systemctl restart evcc.service;'

sync-with-andig-and-deploy:
	# show commits in browser
	open https://github.com/evcc-io/evcc/commits/master
	@echo "Proceed with build? (Y|n)"
	@read line; if [ "$$line" != "Y" ]; then echo aborting; exit 1 ; fi
	@echo building...

	# following may show error if remote upstream alread configured
	# git remote add upstream https://github.com/andig/evcc.git
	# show github remote streams
	git remote -v
	# get new changes
	git fetch upstream
	# local branch
	git checkout master
	# get changes
	git merge upstream/master --no-edit --no-commit

	go mod vendor

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

# patch asn1.go to allow Elli buggy certificates to be accepted with EEBUS
patch-asn1-sudo:
	# echo $(GOROOT)
	cat $(GOROOT)/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go | grep -C 1 "out = true"
	sudo patch -N -t -d $(GOROOT)/src/vendor/golang.org/x/crypto/cryptobyte -i $(CURRDIR)/patch/asn1.diff
	cat $(GOROOT)/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go | grep -C 1 "out = true"

patch-asn1::
	# echo $(GOROOT)
	cat $(GOROOT)/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go | grep -C 1 "out = true"
	patch -N -t -d $(GOROOT)/src/vendor/golang.org/x/crypto/cryptobyte -i $(CURRDIR)/patch/asn1.diff
	cat $(GOROOT)/src/vendor/golang.org/x/crypto/cryptobyte/asn1.go | grep -C 1 "out = true"
