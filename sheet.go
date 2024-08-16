package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

const topBorder = "----------------------------------------\n"

type modelSheet struct {
	data character
}

func (m *modelSheet) Init() tea.Cmd {
	return nil
}
func (m *modelSheet) View() string {
    s := fmt.Sprintf("name: %s\n", m.data.Name)
    s += fmt.Sprintf("level: %d\n", m.data.Level)
    s += fmt.Sprintf("Class: %s\n", m.data.Class)
    s += fmt.Sprintf("Race: %s\n", m.data.Race)
    s += "\n"
	s += topRow(m.data.Attributes, m.data.Savings)
	return s
}

/*
"paralyze": 13,
"petrification": 14,
"staff": 15,
"breath": 16,
"spells": 16
*/

func topRow(attr map[string]int, savings map[string]int) string {

	header := fmt.Sprintf("%s%16s\n", "Attributes", "Savings")
    header += topBorder
	result := header
    result += fmt.Sprintf("%s%4d%20s%4d%8s\n", "str", attr["str"], "paralyze", savings["paralyze"], "|")
	result += fmt.Sprintf("%s%4d%20s%4d%8s\n", "int", attr["int"], "petrification", savings["petrification"],"|")
	result += fmt.Sprintf("%s%4d%20s%4d%8s\n", "wis", attr["wis"], "staff", savings["staff"],"|")
	result += fmt.Sprintf("%s%4d%20s%4d%8s\n", "dex", attr["dex"], "breath", savings["breath"],"|")
	result += fmt.Sprintf("%s%4d%20s%4d%8s\n", "con", attr["con"], "paralyze", savings["paralyze"],"|")
	result += fmt.Sprintf("%s%4d%20s%4d%8s\n", "cha", attr["cha"], "paralyze", savings["paralyze"],"|")
	return result
}
func (m *modelSheet) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter", " ":
		}
	}
	return m, nil

}
