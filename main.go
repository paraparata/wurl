package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

var openapiPathFlag = flag.String("file", "", "List paths from an openapi file")

func main() {
	flag.Parse()
	file, err := os.ReadFile(*openapiPathFlag)
	if err != nil {
		log.Fatal("Error occured")
	}

	data := NewOpenapi(file)
	dataTitle, _, _ := data.Info()
	title := "🚧wurl | " + dataTitle

	uiModel := NewEndpointList(title, data.EndpointsLen(), data.Endpoints())
	p := tea.NewProgram(uiModel, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
