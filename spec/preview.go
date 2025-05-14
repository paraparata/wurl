package spec

import (
	"github.com/pb33f/libopenapi/datamodel/high/base"
	"github.com/pb33f/libopenapi/renderer"
)

func PreviewJSON(schema *base.Schema) (string, error) {
  mg := renderer.NewMockGenerator(renderer.JSON)
  mg.SetPretty()

	mock, err := mg.GenerateMock(schema, "")
	if err != nil {
		return "nil", err
	}
	return string(mock), nil
}
