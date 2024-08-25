package openapi

import (
	"fmt"

	"github.com/pb33f/libopenapi"
	v3 "github.com/pb33f/libopenapi/datamodel/high/v3"
)

type Endpoint struct {
	ID     string
	Method string
	Desc   string
	Path   string
}

type OpenApi struct {
	model     *v3.Document
	endpoints []Endpoint
}

func NewV3(store *[]byte) *OpenApi {
	document, err := libopenapi.NewDocument(*store)
	if err != nil {
		panic(fmt.Sprintf("cannot create new document: %e", err))
	}

	docModel, errors := document.BuildV3Model()
	if len(errors) > 0 {
		for i := range errors {
			fmt.Printf("error: %e\n", errors[i])
		}
		panic(fmt.Sprintf("cannot create v3 model from document: %d errors reported", len(errors)))
	}

	o := &OpenApi{
		model: &docModel.Model,
	}

	itemsLen := 0
	for pathPairs := o.model.Paths.PathItems.First(); pathPairs != nil; pathPairs = pathPairs.Next() {
		pathItem := pathPairs.Value()
		for operations := pathItem.GetOperations().First(); operations != nil; operations = operations.Next() {
			itemsLen++
		}
	}

	i := 0
	o.endpoints = make([]Endpoint, itemsLen)
	for pathPairs := o.model.Paths.PathItems.First(); pathPairs != nil; pathPairs = pathPairs.Next() {
		pathName := pathPairs.Key()
		pathItem := pathPairs.Value()

		for operations := pathItem.GetOperations().First(); operations != nil; operations = operations.Next() {
			method := operations.Key()
			desc := "N/A"

			if operations.Value().Description != "" {
				desc = operations.Value().Description
			}

			o.endpoints[i] = Endpoint{operations.Value().OperationId, method, desc, pathName}
			i += 1
		}
	}

	return o
}

func (o *OpenApi) GetEndpoints() *[]Endpoint { return &o.endpoints }
