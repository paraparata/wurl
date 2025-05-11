package main

import (
	"fmt"

	"github.com/pb33f/libopenapi"
	"github.com/pb33f/libopenapi/datamodel/high/base"
	v3 "github.com/pb33f/libopenapi/datamodel/high/v3"
)

type Endpoint struct {
	path      string
	method    string
	operation *v3.Operation
}

type PayloadProperty struct {
	// TODO: Add extensions and examples
	// examples    []string
	propTypes   []string
	description string
	name        string
}

type Payload struct {
	properties []PayloadProperty
	format     string
}

func (e *Endpoint) RequestBodyPayload() []Payload {
	payloads := make([]Payload, 0)

	if e.operation.RequestBody == nil {
		return payloads
	}

	for req := e.operation.RequestBody.Content.First(); req != nil; req = req.Next() {
		payload := Payload{format: req.Key()}

		for n := req.Value().Schema.Schema().Properties.First(); n != nil; n = n.Next() {
			if n.Value().IsReference() {
				payload.properties = append(payload.properties, PayloadProperty{name: n.Key()})
			} else {
				payload.properties = append(payload.properties, PayloadProperty{
					// TODO: Add extensions and examples
					propTypes:   n.Value().Schema().Type,
					description: n.Value().Schema().Description,
					name:        n.Key(),
				})
			}
		}
		payloads = append(payloads, payload)
	}

	return payloads
}

type Openapi struct {
	docModel     *libopenapi.DocumentModel[v3.Document]
	endpoints    []Endpoint
	endpointsLen int
}

func NewOpenapi(file []byte) *Openapi {
	document, errDoc := libopenapi.NewDocument(file)
	if errDoc != nil {
		panic(fmt.Sprintf("cannot create new document: %e", errDoc))
	}

	docModel, errModel := document.BuildV3Model()
	if len(errModel) > 0 {
		for i := range errModel {
			fmt.Printf("error: %e\n", errModel[i])
		}
		panic(fmt.Sprintf("cannot create v3 model from document: %d errors reported", len(errModel)))
	}

	endpointsLen := 0
	for item := docModel.Model.Paths.PathItems.First(); item != nil; item = item.Next() {
		endpointsLen += item.Value().GetOperations().Len()
	}

	i := 0
	endpoints := make([]Endpoint, endpointsLen)
	for item := docModel.Model.Paths.PathItems.First(); item != nil; item = item.Next() {
		path := item.Key()
		for operation := item.Value().GetOperations().First(); operation != nil; operation = operation.Next() {
			endpoints[i] = Endpoint{
				path:      path,
				method:    operation.Key(),
				operation: operation.Value(),
			}
			i++
		}
	}

	return &Openapi{docModel, endpoints, endpointsLen}
}

func (o *Openapi) Endpoints() []Endpoint {
	return o.endpoints
}

func (o *Openapi) EndpointsLen() int {
	return o.endpointsLen
}

func (o *Openapi) Info() *base.Info {
	return o.docModel.Model.Info
}
