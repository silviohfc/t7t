package main

import (
	"fmt"
	"os"

	"t7t/internal/model"
	"t7t/internal/ui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	store, err := model.NewStore()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao inicializar armazenamento: %v\n", err)
		os.Exit(1)
	}

	app := ui.NewApp(store)

	p := tea.NewProgram(app, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao executar aplicacao: %v\n", err)
		os.Exit(1)
	}
}
