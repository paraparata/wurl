package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/paraparata/wurl2/spec"
	"github.com/paraparata/wurl2/wurl"
)

var openapiPathFlag = flag.String("file", "", "List paths from an openapi file")

func main() {
	flag.Parse()
	file, err := os.ReadFile(*openapiPathFlag)
	if err != nil {
		log.Fatal("Error occured")
	}

	api := spec.New(file)
	wurl := wurl.New(api)
	p := tea.NewProgram(wurl, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
