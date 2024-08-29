package main

import (
	"fmt"
	"github.com/pb33f/libopenapi"
	"github.com/pb33f/libopenapi/renderer"
	"os"
)

func renderFries() string {

	// create a new JSON mock generator
	mg := renderer.NewMockGenerator(renderer.JSON)

	// tell the mock generator to pretty print the output
	mg.SetPretty()

	burgerShop, _ := os.ReadFile("burgershop.openapi.yaml")

	// create a new document from specification and build a v3 model.
	document, _ := libopenapi.NewDocument(burgerShop)
	v3Model, _ := document.BuildV3Model()

	// create a mock of the Fries model
	friesModel := v3Model.Model.Paths.PathItems.First().Value().Options.RequestBody.Content.First().Value().Schema

	// build the fries schema
	fries := friesModel.Schema()

	// generate a mock of the fries schema
	mock, err := mg.GenerateMock(fries, "")

	if err != nil {
		panic(err)
	}

	// print the mock to stdout
	return string(mock)
}

func main() {
	fmt.Println(renderFries())
}
