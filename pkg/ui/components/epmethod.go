package components

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

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

func EpMethod(method string) string {
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
