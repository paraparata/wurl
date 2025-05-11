package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func (i Endpoint) Title() string {
	return fmt.Sprintf("%s %s", strings.ToUpper(i.method), i.path)
}
func (i Endpoint) Description() string {
	if i.operation.Description == "" {
		return "<no desc>"
	}
	return i.operation.Description
}
func (i Endpoint) FilterValue() string {
	return i.method + " " + i.path
}

func NewEndpointList(title string, itemsLen int, items []Endpoint) list.Model {
	listItems := make([]list.Item, itemsLen)
	for i, item := range items {
		listItems[i] = item
	}
	m := list.New(listItems, newItemDelegate(), 0, 0)
	m.Title = title

	return m
}

func updateEndpointList(msg tea.Msg, m wurlModel) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "enter":
			if i, ok := m.list.SelectedItem().(Endpoint); ok {
				m.choice = &i
				m.current = detailView
			}
			return m, nil
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func viewEndpointList(m wurlModel) string {
	return m.list.View()
}
