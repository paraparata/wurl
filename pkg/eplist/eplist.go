package eplist

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/paraparata/wurl/pkg/config"
	"github.com/paraparata/wurl/pkg/openapi"
	"github.com/paraparata/wurl/pkg/ui/components"
)

type EndpointListLoadedMsg struct{}
type EndpointItemMsg struct {
	Message Item
}
type errorMsg struct {
	Message string
}

func errorCmd(msg string) tea.Cmd {
	return func() tea.Msg {
		return errorMsg{Message: msg}
	}
}

func (m *Model) createList() tea.Cmd {
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

type Model struct {
	cfg       *config.Config
	list      list.Model
	onStartup bool
}

func New(cfg *config.Config) *Model {
	items := make([]list.Item, 6)
	delegate := newItemDelegate(newDelegateKeyMap())
	return &Model{
		cfg:       cfg,
		list:      list.New(items, delegate, 0, 0),
		onStartup: true,
	}
}

func (m *Model) Init() tea.Cmd {
	return m.list.StartSpinner()
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.onStartup {
		var cmds []tea.Cmd

		store, err := os.ReadFile(m.cfg.PathFile)
		cmds = append(cmds, errorCmd(fmt.Sprintf("Error reading file: %e", err)))

		m.cfg.Store = store
		createdListCmd := m.createList()
		cmds = append(cmds, createdListCmd)

		return m, tea.Batch(cmds...)
	}

	var cmds []tea.Cmd

	switch msg.(type) {
	case EndpointListLoadedMsg:
		m.list.StopSpinner()
	case EndpointItemMsg:
		cmds = append(cmds, func() tea.Msg {
			return msg
		})
	}

	updatedList, updatedCmd := m.list.Update(msg)
	m.list = updatedList
	cmds = append(cmds, updatedCmd)
	return m, tea.Batch(cmds...)
}

func (m *Model) View() string {
	on_stratup_str := "false"
	if m.onStartup {
		on_stratup_str = "true"
	}
	return on_stratup_str + "\n" + m.cfg.PathFile + "\n" + m.list.View()
}

type Item struct {
	ep          *openapi.Endpoint
	title, desc string
}

func (i Item) Title() string               { return i.title }
func (i Item) Description() string         { return i.desc }
func (i Item) Endpoint() *openapi.Endpoint { return i.ep }
func (i Item) FilterValue() string         { return i.title }

func newItem(ep *openapi.Endpoint, title, desc string) Item {
	return Item{ep, title, desc}
}

func newItemDelegate(keys *delegateKeyMap) list.DefaultDelegate {
	d := list.NewDefaultDelegate()

	d.UpdateFunc = func(msg tea.Msg, m *list.Model) tea.Cmd {
		var selectedItem Item

		if item, ok := m.SelectedItem().(Item); ok {
			selectedItem = item
		} else {
			return nil
		}

		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, keys.choose):
				return func() tea.Msg {
					return EndpointItemMsg{selectedItem}
				}
			}
		}

		return nil
	}

	return d
}

type delegateKeyMap struct {
	choose key.Binding
}

func newDelegateKeyMap() *delegateKeyMap {
	return &delegateKeyMap{
		choose: key.NewBinding(
			key.WithHelp("o", "choose"),
			key.WithHelp("enter", "choose"),
		),
	}
}
