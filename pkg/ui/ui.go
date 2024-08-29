package ui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/paraparata/wurl/pkg/config"
	"github.com/paraparata/wurl/pkg/eplist"
	"github.com/paraparata/wurl/pkg/openapi"
	"github.com/paraparata/wurl/pkg/ui/components"
)

var uiStyle = lipgloss.NewStyle().Margin(1, 2)

type view int

const (
	listView view = iota
	schemaView
)

func (m *model) createList() tea.Cmd {
	api := openapi.NewV3(&m.cfg.Store)
	endpoints := api.GetEndpoints()
	for i, ep := range *endpoints {
		item := newItem(
			&ep,
			fmt.Sprintf("%s %s", components.EpMethod(ep.Method), ep.Path),
			ep.Desc,
		)
		m.list.InsertItem(i, item)
	}
	m.onStartup = false
	return func() tea.Msg {
		return EndpointListLoadedMsg{}
	}
}

type model struct {
	cfg        *config.Config
	list       list.Model
	onStartup  bool
	schema     components.SchemaModel
	activeView view
}

func New(cfg *config.Config) *model {
	m := &model{
		cfg:        cfg,
		activeView: listView,
	}
	m.list = *eplist.New(m.cfg)

	return m
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case eplist.EndpointItemMsg:
		m.schema = *components.NewSchema(
			msg.Message.Title(),
			msg.Message.Description(),
			msg.Message.Endpoint())
		m.activeView = schemaView
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "esc":
			m.activeView = listView
			// m.list, cmd = m.list.Update(msg)
			return m, nil
		}
	}

	return m, cmd
}

func (m model) View() string {
	if m.activeView == 0 {
		return uiStyle.Render(m.list.View())
	}
	return uiStyle.Render(m.schema.View())
}
