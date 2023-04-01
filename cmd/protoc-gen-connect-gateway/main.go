package main

import (
	"fmt"
	"os"

	"go.vallahaye.net/connect-gateway"
	"go.vallahaye.net/connect-gateway/cmd/protoc-gen-connect-gateway/internal/gengateway"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

const usage = "See https://go.vallahaye.net/connect-gateway to learn how to use this plugin.\n\nFlags:\n  -h, --help\tPrint this help and exit.\n      --version\tPrint the version and exit."

func main() {
	if len(os.Args) == 2 && os.Args[1] == "--version" {
		fmt.Fprintln(os.Stdout, connectgateway.Version)
		os.Exit(0)
	}
	if len(os.Args) == 2 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		fmt.Fprintln(os.Stdout, usage)
		os.Exit(0)
	}
	if len(os.Args) != 1 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}
	protogen.Options{}.Run(
		func(plugin *protogen.Plugin) error {
			plugin.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
			for _, file := range plugin.Files {
				if file.Generate {
					gengateway.Generate(plugin, file)
				}
			}
			return nil
		},
	)
}
