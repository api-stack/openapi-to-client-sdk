package generators

import (
	"fmt"
	"log"
	"os"

	"github.com/getkin/kin-openapi/openapi3"
)

func GenerateClientSDK(lang string, openApiSpec string) error {
	log.Printf("Generating client sdk for spec %s", openApiSpec)
	// Reading open api spec
	spec, err := os.ReadFile(openApiSpec)
	if err != nil {
		panic(fmt.Sprintf("Error while reading open api spec : %s", err.Error()))
	}
	loader := openapi3.NewLoader()
	doc, err := loader.LoadFromData(spec)
	if err != nil {
		panic(fmt.Sprintf("Error while loading openapi spec : %s", err.Error()))
	}
	err = doc.Validate(loader.Context)
	if err != nil {
		panic(fmt.Sprintf("Error while loading openapi spec : %s", err.Error()))
	}

	return nil
}
