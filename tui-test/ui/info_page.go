package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type InfoPageModel struct {
	message string
}

func NewInfoPageModel() InfoPageModel {
	return InfoPageModel{
		message: "This is a simple info page.\nPress ESC to return.",
	}
}

func NewInfoPageModelWithText(text string) InfoPageModel {
	return InfoPageModel{
		message: fmt.Sprintf("You typed:\n\n%s\n\nPress ESC to return.", text),
	}
}

func (m InfoPageModel) Init() tea.Cmd {
	return nil
}

func (m InfoPageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return NewMenuModel(), nil
		case "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m InfoPageModel) View() string {
	return m.message
}
