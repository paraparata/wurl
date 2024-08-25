package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/paraparata/wurl/pkg/config"
	"github.com/paraparata/wurl/pkg/openapi"
	"github.com/paraparata/wurl/pkg/ui/components"
)

var uiStyle = lipgloss.NewStyle().Margin(1, 2)

type view int

const (
	listView view = iota
	schemaView
)

type model struct {
	*config.Config
	list       list.Model
	schema     components.SchemaModel
	activeView view
}

func New(cfg *config.Config) *model {
	m := &model{
		Config:     cfg,
		activeView: listView,
	}

	api := openapi.NewV3(m.Store)
	endpoints := api.GetEndpoints()
	items := make([]list.Item, len(*endpoints))
	for i, ep := range *endpoints {
		items[i] = components.NewEpListItem(
			&ep,
			fmt.Sprintf("%s %s", components.EpMethod(ep.Method), ep.Path),
			ep.Desc,
		)
	}

	delegate := components.NewEpListItemDelegate(components.NewDelegateEpListKeyMap())
	operations := list.New(items, delegate, 0, 0)
	operations.Title = "wurl"
	m.list = operations

	return m
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "esc":
			m.activeView = listView
			// m.list, cmd = m.list.Update(msg)
			return m, nil
		case "enter":
			ep := m.list.SelectedItem().(components.EpListItem).Endpoint()
			m.schema = *components.NewSchema(ep.Method, ep.Path)
			m.activeView = schemaView
			return m, nil
		}
	case tea.WindowSizeMsg:
		h, v := uiStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.activeView == 0 {
		return uiStyle.Render(m.list.View())
	}
	return uiStyle.Render(m.schema.View())
}
