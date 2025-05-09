package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/pb33f/libopenapi"
	v3 "github.com/pb33f/libopenapi/datamodel/high/v3"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

var (
	get = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#000000")).
		Background(lipgloss.Color("#61affe"))
	pos = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#000000")).
		Background(lipgloss.Color("#49cc90"))
	put = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#000000")).
		Background(lipgloss.Color("#F18F01"))
	del = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#000000")).
		Background(lipgloss.Color("#f93e3e"))
)

func EpMethodStyle(method string) string {
	mtd := fmt.Sprintf("[%s]", strings.ToUpper(method))
	switch method {
	case "get":
		return get.Render(mtd)
	case "post":
		return pos.Render(mtd)
	case "put", "patch":
		return put.Render(mtd)
	case "delete":
		return del.Render(mtd)
	}

	return mtd
}

type Endpoint struct {
	path      string
	method    string
	operation *v3.Operation
}

func (i Endpoint) Title() string {
	return fmt.Sprintf("%s %s", EpMethodStyle(i.method), i.path)
}
func (i Endpoint) Description() string {
	if i.operation.Description == "" {
		return "<no desc>"
	}
	return i.operation.Description
}
func (i Endpoint) FilterValue() string {
	return fmt.Sprintf("%s %s", i.method, i.path)
}

type model struct {
	list list.Model
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

var openapiPathFlag = flag.String("openapi", "", "List paths from an openapi file")

func main() {
	flag.Parse()
	file, _ := os.ReadFile(*openapiPathFlag)
	document, errDoc := libopenapi.NewDocument(file)
	if errDoc != nil {
		panic(fmt.Sprintf("cannot create new document: %e", errDoc))
	}

	docModel, errModel := document.BuildV3Model()
	if len(errModel) > 0 {
		for i := range errModel {
			fmt.Printf("error: %e\n", errModel[i])
		}
		panic(fmt.Sprintf("cannot create v3 model from document: %d errors reported", len(errModel)))
	}

	endpointsLen := 0
	for item := docModel.Model.Paths.PathItems.First(); item != nil; item = item.Next() {
		endpointsLen += item.Value().GetOperations().Len()
	}

	i := 0
	endpoints := make([]Endpoint, endpointsLen)
	for item := docModel.Model.Paths.PathItems.First(); item != nil; item = item.Next() {
		path := item.Key()
		for operation := item.Value().GetOperations().First(); operation != nil; operation = operation.Next() {
			endpoints[i] = Endpoint{
				path:      path,
				method:    operation.Key(),
				operation: operation.Value(),
			}
			i++
		}
	}

	listItems := make([]list.Item, endpointsLen)
	for i, endpoint := range endpoints {
		listItems[i] = endpoint
	}

	m := model{list: list.New(listItems, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "🚧wurl | " + docModel.Model.Info.Title

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
