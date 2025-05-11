package main

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type view int

const (
	listView view = iota
	detailView
)

type wurlModel struct {
	list    list.Model
	choice  *Endpoint
	current view
}

func NewWurl(api *Openapi) *wurlModel {
	title := "🚧wurl | " + api.Info().Title

	listItems := make([]list.Item, api.EndpointsLen())
	for i, item := range api.Endpoints() {
		listItems[i] = item
	}
	l := list.New(listItems, newItemDelegate(), 0, 0)
	l.Title = title

	m := wurlModel{list: l}

	return &m
}

func (m wurlModel) Init() tea.Cmd {
	return m.list.StartSpinner()
}

func (m wurlModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if keypress := msg.String(); keypress == "ctrl+c" {
			return m, tea.Quit
		}
	}

	if m.current == listView {
		return updateEndpointList(msg, m)
	}
	return updateEndpointDetail(msg, m)
}

func (m wurlModel) View() string {
	if m.current == detailView {
		return docStyle.Render(viewEndpointDetail(m))
	}
	return docStyle.Render(viewEndpointList(m))
}
