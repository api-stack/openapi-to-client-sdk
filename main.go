package main

import (
	"os"

	"github.com/api-stack/openapi-to-client-sdk/generators"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		panic("Usage: apistack generate node|go|android|* <path-to-openapi-spec>")
	}
	switch args[0] {
	case "generate":
		lang := args[1]
		err := generators.GenerateClientSDK(lang, args[2])
		if err != nil {
			panic(err)
		}
	default:
		panic("Invalid command")
	}
}
