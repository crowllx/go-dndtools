package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

type sessionModel struct{
    events []string
}

func (m *sessionModel) Init() tea.Cmd {
	return nil
}
func (m *sessionModel) View() string {
    var s string
    for _, e := range m.events {
        s += fmt.Sprintf("  %s\n", e)
    }
	return s
}
func (m *sessionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}
