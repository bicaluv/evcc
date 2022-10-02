# SunSpec tools for Go

This package contains Go data types representing the [SunSpec][] information model.
Subpackages implement particular use-cases that rely on this domain model:

 * [/impl](./impl) contains the types used by the remainder of the library
 * [/smdx](./smdx) contains types that represent (XML) SMDX models defined by the Sunspec specification
 * [/generators](./generators) implements code generators that transform the [SMDX models][SMDX] into Go code
 * [/models](./models) contains one generated package for each SMDX model
 * [/xml](./xml) contains utilities for exchanging SunSpec data using XML format
 * [/modbus](./modbus) contains utilities for talking to SunSpec devices via Modbus TCP

[SunSpec]: http://sunspec.org/
[SMDX]: https://github.com/sunspec/models

This implementation is based on the original work from [crabmusket/gosunspec](https://github.com/crabmusket/gosunspec/issues/38).

## Generated code

This package uses typesafe representations of each model type. To avoid
excessive manual maintenance, structs for each type of model defined by SunSpec
are generated from the [SMDX model files][SMDX]. To regenerate this code, first
initialise the spec submodule with:

    git submodule update --init spec

And from the `spec` directory:

    git pull origin master

Then run the generators:

    go generate ./models
