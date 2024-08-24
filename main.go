package main

import (
	"flag"
	"fmt"
	"strings"

	// "github.com/paraparata/wurl/ui"
	"os"

	"github.com/pb33f/libopenapi"
)

var openapiPathFlag = flag.String("openapi", "openapi.yml", "List paths from an openapi file")

func main() {
	flag.Parse()

	// load in the petstore sample OpenAPI specification
	// into a byte slice.
	petstore, _ := os.ReadFile(*openapiPathFlag)

	// create a new Document from from the byte slice.
	document, err := libopenapi.NewDocument(petstore)

	// if anything went wrong, an error is thrown
	if err != nil {
		panic(fmt.Sprintf("cannot create new document: %e", err))
	}

	// because we know this is a v3 spec, we can build a ready to go model from it.
	docModel, errors := document.BuildV3Model()

	// if anything went wrong when building the v3 model, a slice of errors will be returned
	if len(errors) > 0 {
		for i := range errors {
			fmt.Printf("error: %e\n", errors[i])
		}
		panic(fmt.Sprintf("cannot create v3 model from document: %d errors reported", len(errors)))
	}

	for pathPairs := docModel.Model.Paths.PathItems.First(); pathPairs != nil; pathPairs = pathPairs.Next() {
		pathName := pathPairs.Key()
		pathItem := pathPairs.Value()
		fmt.Printf("Path %s has %d operations\n", pathName, pathItem.GetOperations().Len())

		for operations := pathItem.GetOperations().First(); operations != nil; operations = operations.Next() {
			name := operations.Key()
			fmt.Printf("> %s\n", strings.ToUpper(name))
		}
	}

	// if _, err := ui.NewProgram().Run(); err != nil {
	// 	fmt.Println("Error running program:", err)
	// 	os.Exit(1)
	// }
}
