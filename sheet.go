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
    s := fmt.Sprintf("name: \t%s\texp:\t%d\n", m.data.Name, m.data.Exp)
    s += fmt.Sprintf("level: \t%d\tcoin:\t%d\n", m.data.Level, m.data.Coin)
    s += fmt.Sprintf("Class: \t%s\n", m.data.Class)
    s += fmt.Sprintf("Race: \t%s\n", m.data.Race)
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

	header := fmt.Sprintf("%s\t%-12s\n", "Attributes", "Savings")
    header += topBorder
	result := header
    result += fmt.Sprintf("%-4s%-4d\t%-14s%4d%8s\n", "str", attr["str"], "paralyze", savings["paralyze"], "|")
	result += fmt.Sprintf("%-4s%-4d\t%-14s%4d%8s\n", "int", attr["int"], "petrification", savings["petrification"],"|")
	result += fmt.Sprintf("%-4s%-4d\t%-14s%4d%8s\n", "wis", attr["wis"], "staff", savings["staff"],"|")
	result += fmt.Sprintf("%-4s%-4d\t%-14s%4d%8s\n", "dex", attr["dex"], "breath", savings["breath"],"|")
	result += fmt.Sprintf("%-4s%-4d\t%-14s%4d%8s\n", "con", attr["con"], "paralyze", savings["paralyze"],"|")
	result += fmt.Sprintf("%-4s%-4d\t%-14s%4d%8s\n", "cha", attr["cha"], "paralyze", savings["paralyze"],"|")
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
