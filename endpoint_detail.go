package main

import (
	// "fmt"

	// "fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func updateEndpointDetail(msg tea.Msg, m wurlModel) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q":
			m.choice = nil
			m.current = listView
			return m, nil
		}
	}

	var cmd tea.Cmd
	return m, cmd
}

func viewEndpointDetail(m wurlModel) string {
	var (
		sections []string
	)

	sections = append(sections, m.choice.Title())
	sections = append(sections, m.choice.Description())
	sections = append(sections, "Parameters")
	for _, param := range m.choice.operation.Parameters {
		sections = append(sections, "  - "+param.Name)
		sections = append(sections, "    "+param.In)
		sections = append(sections, "    "+lipgloss.
			NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#C2B8C2", Dark: "#4D4D4D"}).
			Render(param.Description))
	}
	sections = append(sections, "Request Body")
	if m.choice.operation.RequestBody != nil {
		sections = append(sections, m.choice.operation.RequestBody.Description)
		for _, payload := range m.choice.RequestBodyPayload() {
			sections = append(sections, "  - "+payload.format)
			for _, prop := range payload.properties {
				sections = append(sections, "     - "+prop.name)
				sections = append(sections, "       "+prop.description)
				// sections = append(sections, "    "+prop.)
			}
		}
	}

	return lipgloss.JoinVertical(lipgloss.Left, sections...)
	// return lipgloss.NewStyle().Margin(1, 0, 2, 4).Render(fmt.Sprintf("%s? Sounds good to me.", m.choice.Title()))
}
