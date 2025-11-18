package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/list"
)

type menuItem string

func (i menuItem) FilterValue() string { return "" }

type MenuModel struct {
	list list.Model
}

func NewMenuModel() MenuModel {
	items := []list.Item{
		menuItem("Text Input Page"),
		menuItem("Information Page"),
		menuItem("Quit"),
	}

	l := list.New(items, list.NewDefaultDelegate(), 30, 10)
	l.Title = "Main Menu"

	return MenuModel{list: l}
}

func (m MenuModel) Init() tea.Cmd {
	return nil
}

func (m MenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "enter":
			selected := m.list.SelectedItem().(menuItem)

			switch selected {
			case "Text Input Page":
				return NewInputPageModel(), nil
			case "Information Page":
				return NewInfoPageModel(), nil
			case "Quit":
				return m, tea.Quit
			}

		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m MenuModel) View() string {
	return m.list.View()
}
