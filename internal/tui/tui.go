package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

// func () {

// }

type Model struct {
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m Model) View() string {
	s := "FASTCOPY"
	return s
}

func initialModel() Model {
	return Model{}
}

func NewTui() *tea.Program {
	return tea.NewProgram(initialModel(), tea.WithAltScreen())
}
