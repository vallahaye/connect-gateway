# Connect-Gateway

[![PkgGoDev](https://pkg.go.dev/badge/go.vallahaye.net/connect-gateway)](https://pkg.go.dev/go.vallahaye.net/connect-gateway) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/vallahaye/connect-gateway) [![GoReportCard](https://goreportcard.com/badge/github.com/vallahaye/connect-gateway)](https://goreportcard.com/badge/github.com/vallahaye/connect-gateway) ![GitHub](https://img.shields.io/github/license/vallahaye/connect-gateway)

The Connect-Gateway introduces direct binding from [gRPC-Gateway](https://grpc-ecosystem.github.io/grpc-gateway/) local handlers to [Connect](https://connect.build/) service handlers. It addresses the recurring request to support Google API HTTP annotations in Connect:

- [Service option for Connect HTTP path #468](https://github.com/bufbuild/connect-go/issues/468)
- [`google.api.Http` annotation support #274](https://github.com/bufbuild/connect-go/issues/274)

We provide a complete solution for the two to communicate seamlessly through simple function calls, without relying on network communications. All of which is done by reusing as much of the code from both projects as possible and mimicking the [connect-go](https://github.com/bufbuild/connect-go) implementation, i.e. reducing code generation as much as possible, with most of the logic being provided in a library.

## Features

- [Connect interceptors](https://connect.build/docs/go/interceptors) support
- Bi-directional [gRPC metadata](https://github.com/grpc/grpc-go/blob/master/Documentation/grpc-metadata.md) transmission
- [Connect errors](https://connect.build/docs/go/errors) to [gRPC errors](https://github.com/grpc/grpc-go/blob/master/Documentation/rpc-errors.md) convertion

## Limitations

- No support for streaming calls as [it is not yet supported by the gRPC-Gateway's "in-process" transport mode](https://github.com/grpc-ecosystem/grpc-gateway/blob/main/protoc-gen-grpc-gateway/internal/gengateway/template.go#L609)
- Uninitialized [Request.Peer](https://pkg.go.dev/github.com/bufbuild/connect-go#Request.Peer) and [Request.Spec](https://pkg.go.dev/github.com/bufbuild/connect-go#Request.Spec) properties on Connect requests as it cannot be set externally

## Example

Please refer to the [example/](https://github.com/vallahaye/connect-gateway/tree/main/example) directory for a basic example of how to integrate the Connect-Gateway into your project.

## Credit

Some codes were copied pasted from the [connect-go](https://github.com/bufbuild/connect-go) code base, in particular:

- The `HandlerOption` system with its interceptor chaining capability
- The `wrapComments` utility function in the code generator
