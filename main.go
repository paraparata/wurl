package main

import (
	"flag"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/paraparata/wurl/pkg/config"
	"github.com/paraparata/wurl/pkg/ui"
)

var openapiPathFlag = flag.String("openapi", "", "List paths from an openapi file")

func main() {
	flag.Parse()

	store, err := os.ReadFile(*openapiPathFlag)
	if err != nil {
		panic(fmt.Sprintf("Error reading file: %e", err))
	}

	cfg := config.New(&store)

	model := ui.New(cfg)
	if _, err := tea.NewProgram(model, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
