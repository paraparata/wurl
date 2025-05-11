package wurl

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/paraparata/wurl2/spec"
)

func newEndpointList(api *spec.Spec) list.Model {
	listItems := make([]list.Item, api.EndpointsLen())
	for i, item := range api.Endpoints() {
		listItems[i] = listItemModel{Endpoint: &item}
	}

	model := list.New(listItems, newItemDelegate(), 0, 0)
	model.Title = "🚧wurl | " + api.Info().Title

	return model
}

func updateEndpointList(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "enter":
			if i, ok := m.list.SelectedItem().(listItemModel); ok {
				m.activeItem = &i
				c := lipgloss.NewStyle().Width(m.width).Render(endpointDetailContent(&i))
				m.viewport.SetContent(c)
				// m.viewport.SetContent(c + "\n" + c + "\n" + c + "\n" + c + "\n" + c + "\n" + c + "\n" + c)
			}
			return m, nil
		}
	case tea.WindowSizeMsg:
		m.list.SetSize(m.width, m.height)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func viewEndpointList(m model) string {
	return m.list.View()
}
