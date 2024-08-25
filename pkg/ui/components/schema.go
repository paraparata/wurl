package components

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type SchemaModel struct {
	method string
	path   string
}

func NewSchema(method, path string) *SchemaModel {
	return &SchemaModel{method, path}

}
func (m SchemaModel) Init() tea.Cmd {
	return nil
}

func (m SchemaModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m SchemaModel) View() string {
	return fmt.Sprintf("%s %s", EpMethod(m.method), m.path)
}
