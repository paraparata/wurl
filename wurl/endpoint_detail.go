package wurl

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/paraparata/wurl/spec"
	"github.com/paraparata/wurl/styles"
)

func newEndpointDetail() viewport.Model {
	model := viewport.New(0, 0)
	model.MouseWheelEnabled = true
	return model
}

func updateEndpointDetail(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q":
			m.activeItem = nil
			m.viewport.GotoTop()
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}

func viewEndpointDetail(m model) string {
	var sections []string

	sections = append(sections, styles.Titleize(styles.Method(m.activeItem.Method), m.activeItem.Path))
	sections = append(sections, "\n"+m.viewport.View())

	// helpView := m.help.View(viewport.KeyMap.Down.Help())
	// sections = append(sections, helpView)
	return lipgloss.NewStyle().Padding(0, 2).Render(
		lipgloss.JoinVertical(lipgloss.Left, sections...),
	)
}

func endpointDetailContent(activeItem *listItemModel) string {
	var sections []string

	desc, errDesc := glamour.Render(activeItem.Description(), "dark")
	if errDesc == nil {
		sections = append(sections, desc)
	}

	if len(activeItem.Operation.Parameters) != 0 {
		sections = append(sections, styles.Heading.Render("Parameters"))
		for _, param := range activeItem.Operation.Parameters {
			paramText := "  - " + param.Name

			required := ""
			if *param.Required {
				required = ", required"
			}

			paramAttr := styles.Dimmed.
				Render(fmt.Sprintf("in: %s%s", param.In, required))
			paramText += " " + styles.Italic(paramAttr)
			sections = append(sections, paramText)
			sections = append(sections, "    "+styles.Dimmed.
				Render(param.Description))
		}
	}

	if activeItem.Operation.RequestBody != nil && len(activeItem.RequestBodyPayload()) != 0 {
		sections = append(sections, styles.Heading.Render("Request Body"))
		sections = append(sections, activeItem.Operation.RequestBody.Description)
		for _, payload := range activeItem.RequestBodyPayload() {
			sections = append(sections, "  - "+payload.Format)
			for _, prop := range payload.Properties {
				sections = append(sections, "     - "+prop.Name)
				sections = append(sections, "       "+styles.Dimmed.Render(prop.Description))
				// sections = append(sections, "    "+prop.)
			}
		}

		preview, _ := spec.PreviewJSON(activeItem.Operation.RequestBody.Content.Newest().Value.Schema.Schema())
		sections = append(sections, styles.CodeViewer(preview))
	}

	sections = append(sections, styles.Heading.Render("Responses"))
	if activeItem.Operation.Responses != nil {
		for resCode := activeItem.Operation.Responses.Codes.First(); resCode != nil; resCode = resCode.Next() {
			statusCode := resCode.Key()
			sections = append(sections, "  - "+statusCode)
			sections = append(sections, "    "+styles.Dimmed.Render(resCode.Value().Description))

			if resCode.Value().Headers != nil && resCode.Value().Headers.Len() != 0 {
				sections = append(sections, "    "+"Headers")
				for header := resCode.Value().Headers.First(); header != nil; header = header.Next() {
					headerName := header.Key()
					sections = append(sections, "     - "+headerName)
					if header.Value().Schema == nil {
						break
					}
					sections = append(sections, "       "+header.Value().Schema.Schema().Type[0])
					sections = append(sections, "       "+strconv.FormatBool(header.Value().Required))
					sections = append(sections, "       "+styles.Dimmed.Render(header.Value().Description))
				}
			}

			if resCode.Value().Content != nil && resCode.Value().Content.Len() != 0 {
				sections = append(sections, "    "+"Response")
				for response := resCode.Value().Content.First(); response != nil; response = response.Next() {
					responseType := response.Key()
					sections = append(sections, "     - "+responseType)
					if response.Value().Schema == nil {
						break
					}

					resType := strings.Join(response.Value().Schema.Schema().Type, "")
					sections = append(sections, "       "+resType)
					sections = append(sections, "       "+styles.Dimmed.Render(response.Value().Schema.Schema().Description))

					preview, _ := spec.PreviewJSON(response.Value().Schema.Schema())
					sections = append(sections, styles.CodeViewer(preview))
				}
			}
		}
	}

	return lipgloss.JoinVertical(lipgloss.Left, sections...)
}
