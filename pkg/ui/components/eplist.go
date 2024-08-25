package components

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/paraparata/wurl/pkg/openapi"
)

type EpListItem struct {
	ep          *openapi.Endpoint
	title, desc string
}

func (i EpListItem) Title() string               { return i.title }
func (i EpListItem) Description() string         { return i.desc }
func (i EpListItem) Endpoint() *openapi.Endpoint { return i.ep }
func (i EpListItem) FilterValue() string         { return i.title }

func NewEpListItem(ep *openapi.Endpoint, title, desc string) EpListItem {
	return EpListItem{ep, title, desc}

}

type delegateKeyMap struct {
	choose key.Binding
}

func NewDelegateEpListKeyMap() *delegateKeyMap {
	return &delegateKeyMap{
		choose: key.NewBinding(
			key.WithKeys("o", "enter"),
			key.WithHelp("o", "choose"),
		),
	}
}

func NewEpListItemDelegate(keys *delegateKeyMap) list.DefaultDelegate {
	d := list.NewDefaultDelegate()

	d.UpdateFunc = func(msg tea.Msg, m *list.Model) tea.Cmd {
		var title string

		if i, ok := m.SelectedItem().(EpListItem); ok {
			title = i.Title()
		} else {
			return nil
		}

		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, keys.choose):
				return m.NewStatusMessage("You chose " + title)
			}
		}

		return nil
	}

	return d
}
