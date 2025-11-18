package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/textinput"
)

type InputPageModel struct {
	input textinput.Model
}

func NewInputPageModel() InputPageModel {
	ti := textinput.New()
	ti.Placeholder = "Type something..."
	ti.Focus()

	return InputPageModel{input: ti}
}

func (m InputPageModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m InputPageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			return NewInfoPageModelWithText(m.input.Value()), nil
		case "esc":
			return NewMenuModel(), nil
		case "q":
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

func (m InputPageModel) View() string {
	return fmt.Sprintf(
		"Enter something and press Enter:\n\n%s\n\n(esc to go back)",
		m.input.View(),
	)
}
