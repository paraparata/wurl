package wurl

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/paraparata/wurl2/spec"
)

type model struct {
	spec       *spec.Spec
	list       list.Model
	viewport   viewport.Model
	activeItem *listItemModel
	width      int
	height     int
}

type listItemModel struct {
	*spec.Endpoint
}

func (i listItemModel) Title() string {
	return fmt.Sprintf("%s %s", strings.ToUpper(i.Endpoint.Method), i.Endpoint.Path)
}
func (i listItemModel) Description() string {
	if i.Endpoint.Operation.Description == "" {
		return "_no desc_"
	}
	return i.Endpoint.Operation.Description
}
func (i listItemModel) FilterValue() string {
	return i.Endpoint.Method + " " + i.Endpoint.Path
}
