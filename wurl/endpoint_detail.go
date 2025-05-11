package wurl

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/paraparata/wurl/styles"
)

func newEndpointDetail() viewport.Model {
	model := viewport.New(0, 0)
	model.MouseWheelEnabled = true
	return model
}

func updateEndpointDetail(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q":
			m.activeItem = nil
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}

func viewEndpointDetail(m model) string {
	var sections []string

	sections = append(sections, styles.Titleize(styles.Method(m.activeItem.Method), m.activeItem.Path))
	sections = append(sections, m.viewport.View())

	return lipgloss.NewStyle().Padding(0, 2).Render(
		lipgloss.JoinVertical(lipgloss.Left, sections...),
	)
}

func endpointDetailContent(activeItem *listItemModel) string {
	var sections []string

	desc, errDesc := glamour.Render(activeItem.Description(), "dark")
	if errDesc == nil {
		sections = append(sections, desc)
	}

	sections = append(sections, "Parameters")
	for _, param := range activeItem.Operation.Parameters {
		sections = append(sections, "  - "+param.Name)
		sections = append(sections, "    "+param.In)
		sections = append(sections, "    "+lipgloss.
			NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#C2B8C2", Dark: "#4D4D4D"}).
			Render(param.Description))
	}
	sections = append(sections, "Request Body")
	if activeItem.Operation.RequestBody != nil {
		sections = append(sections, activeItem.Operation.RequestBody.Description)
		for _, payload := range activeItem.RequestBodyPayload() {
			sections = append(sections, "  - "+payload.Format)
			for _, prop := range payload.Properties {
				sections = append(sections, "     - "+prop.Name)
				sections = append(sections, "       "+prop.Description)
				// sections = append(sections, "    "+prop.)
			}
		}
	}

	return lipgloss.JoinVertical(lipgloss.Left, sections...)
}
