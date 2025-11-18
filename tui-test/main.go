package main

import (
	tea "github.com/charmbracelet/bubbletea"

	"tui-test/ui" // ← change to your module name
)

func main() {
	model := ui.NewMenuModel()

	p := tea.NewProgram(model)
	if err := p.Start(); err != nil {
		panic(err)
	}
}
