package styles

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	DocStyle = lipgloss.NewStyle().Margin(1, 2)
	Get      = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#000000")).
			Background(lipgloss.Color("#61affe"))
	Pos = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#000000")).
		Background(lipgloss.Color("#49cc90"))
	Put = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#000000")).
		Background(lipgloss.Color("#F18F01"))
	Del = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#000000")).
		Background(lipgloss.Color("#f93e3e"))
	Heading = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#DE76F1"))
	SelectedMethod = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#000000")).
			Background(lipgloss.AdaptiveColor{Light: "#F793FF", Dark: "#AD58B4"})
	DimmedInverse = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#000000")).
			Background(lipgloss.AdaptiveColor{Light: "#A49FA5", Dark: "#777777"})
	Dimmed = lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#C2B8C2", Dark: "#4D4D4D"})
)

func Method(method string) string {
	mtd := fmt.Sprintf(" %s ", strings.ToUpper(method))
	switch method {
	case "get":
		return Get.Render(mtd)
	case "post":
		return Pos.Render(mtd)
	case "put", "patch":
		return Put.Render(mtd)
	case "delete":
		return Del.Render(mtd)
	}

	return mtd
}

// TODO: Styleize the method
func Titleize(method, path string) string {
	return method + " " + path
}

func CodeViewer(text string) string {
	return lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		Foreground(lipgloss.Color("230")).
		Padding(0, 4).Margin(1, 0).Render(text)
}

func Italic(text string) string {
	return lipgloss.NewStyle().Italic(true).Render(text)
}
