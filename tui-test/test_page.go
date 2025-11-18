package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type TestPageModel struct {
	content string
}

func NewTestPage(content string) TestPageModel {
	return TestPageModel{
		content: content,
	}
}

func (m TestPageModel) Init() tea.Cmd {
	return nil
}

func (m TestPageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m TestPageModel) View() string {
	length := len(m.content)

	top := "+" + strings.Repeat("-", length+2) + "+\n"
	mid := fmt.Sprintf("| %s |\n", m.content)
	bottom := "+" + strings.Repeat("-", length+2) + "+\n"

	return top + mid + bottom + "\n(press q to quit)\n"
}
