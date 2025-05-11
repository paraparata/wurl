package wurl

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/paraparata/wurl/spec"
	"github.com/paraparata/wurl/styles"
)

func New(api *spec.Spec) *model {
	m := model{spec: api}
	m.list = newEndpointList(api)
	m.viewport = newEndpointDetail()

	return &m
}

func (m model) Init() tea.Cmd {
	return tea.Batch(tea.SetWindowTitle(m.list.Title))
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if keypress := msg.String(); keypress == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := styles.DocStyle.GetFrameSize()
		m.width = msg.Width - h
		m.height = msg.Height - v
		m.viewport.Width = m.width - h
		m.viewport.Height = m.height - v
	}

	if m.activeItem != nil {
		newModel, cmd := updateEndpointDetail(msg, m)
		cmds = append(cmds, cmd)
		return newModel, tea.Batch(cmds...)
	}
	newModel, cmd := updateEndpointList(msg, m)
	cmds = append(cmds, cmd)
	return newModel, tea.Batch(cmds...)
}

func (m model) View() string {
	if m.activeItem != nil {
		return styles.DocStyle.Render(viewEndpointDetail(m))
	}
	return styles.DocStyle.Render(viewEndpointList(m))
}
