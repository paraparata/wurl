package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

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

type model struct {
	list list.Model
}

func NewEndpointList(title string, itemsLen int, items []Endpoint) *model {
	listItems := make([]list.Item, itemsLen)
	for i, item := range items {
		listItems[i] = item
	}
	m := model{list: list.New(listItems, newItemDelegate(), 0, 0)}
	m.list.Title = title

	return &m
}

func (m model) Init() tea.Cmd {
	return m.list.StartSpinner()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return docStyle.Render(m.list.View())
}
