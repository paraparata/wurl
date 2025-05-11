package spec

import (
	"fmt"

	"github.com/pb33f/libopenapi"
	"github.com/pb33f/libopenapi/datamodel/high/base"
	v3 "github.com/pb33f/libopenapi/datamodel/high/v3"
)

type Endpoint struct {
	Path      string
	Method    string
	Operation *v3.Operation
}

type PayloadProperty struct {
	// TODO: Add extensions and examples
	// examples    []string
	PropTypes   []string
	Description string
	Name        string
}

type Payload struct {
	Properties []PayloadProperty
	Format     string
}

func (e *Endpoint) RequestBodyPayload() []Payload {
	payloads := make([]Payload, 0)

	if e.Operation.RequestBody == nil {
		return payloads
	}

	for req := e.Operation.RequestBody.Content.First(); req != nil; req = req.Next() {
		payload := Payload{Format: req.Key()}

		for n := req.Value().Schema.Schema().Properties.First(); n != nil; n = n.Next() {
			if n.Value().IsReference() {
				payload.Properties = append(payload.Properties, PayloadProperty{Name: n.Key()})
			} else {
				payload.Properties = append(payload.Properties, PayloadProperty{
					// TODO: Add extensions and examples
					PropTypes:   n.Value().Schema().Type,
					Description: n.Value().Schema().Description,
					Name:        n.Key(),
				})
			}
		}
		payloads = append(payloads, payload)
	}

	return payloads
}

type Spec struct {
	docModel     *libopenapi.DocumentModel[v3.Document]
	endpoints    []Endpoint
	endpointsLen int
}

func New(file []byte) *Spec {
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
				Path:      path,
				Method:    operation.Key(),
				Operation: operation.Value(),
			}
			i++
		}
	}

	return &Spec{docModel, endpoints, endpointsLen}
}

func (o *Spec) Endpoints() []Endpoint {
	return o.endpoints
}

func (o *Spec) EndpointsLen() int {
	return o.endpointsLen
}

func (o *Spec) Info() *base.Info {
	return o.docModel.Model.Info
}
