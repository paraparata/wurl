package components

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/paraparata/wurl/pkg/openapi"
)

var (
	layout = lipgloss.NewStyle()
)

type SchemaModel struct {
	ep    *openapi.Endpoint
	title string
	desc  string
}

func NewSchema(title, desc string, ep *openapi.Endpoint) *SchemaModel {
	return &SchemaModel{ep, title, desc}

}
func (m SchemaModel) Init() tea.Cmd {
	return nil
}

func (m SchemaModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m SchemaModel) View() string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		m.title,
		m.desc,
		"\n\nParams:",
		m.listParam(),
		"\n\nRequest Body:",
		m.listReqBody(),
		"\n\nResponse:",
		m.listResBody(),
	)
}

func (m *SchemaModel) listParam() string {
	var paramsStr string
	for _, p := range m.ep.Operation.Parameters {
		paramsStr += fmt.Sprintf("- [%s] (%s) %s\n", p.In, p.Schema.Schema().Type, p.Name)
	}
	return paramsStr
}

func (m *SchemaModel) listReqBody() string {
	return m.ep.GetReqExample()
}

func (m *SchemaModel) listResBody() string {
	return m.ep.GetResExample()
}
