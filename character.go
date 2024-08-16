package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// character data
type clist struct {
	Characters []character `json:"characters,[]characters"`
}
type character struct {
	Name       string         `json:"name"`
	Level      int            `json:"level"`
	Class      string         `json:"class"`
	Race       string         `json:"race"`
	Coin       int            `json:"coin"`
	Exp        int            `json:"exp"`
	Attributes map[string]int `json:"attributes"`
	Savings    map[string]int `json:"savings"`
}

const topBorder = "----------------------------------------\n"

// model to render character sheet
type characterModel struct {
	data          character
	width, height int
}

func (m *characterModel) Init() tea.Cmd {
	return nil
}
func (m *characterModel) View() string {
	s := header(&m.data)
	s += topRow(m.data.Attributes, m.data.Savings)
	lines := strings.Count(s, "\n")
	if lines < m.height-1 {
		s += strings.Repeat("\n", m.height-2-lines)
	}
	subs := strings.Split(s, "\n")
	for i, sub := range subs {
		subs[i] = fmt.Sprintf("%-80s|", sub)
	}
	s = strings.Join(subs, "\n")
	return s
}

func header(c *character) string {
	var s string
	s += fmt.Sprintf("%-8s %-10s %-8s %-8d\n", "name:", c.Name, "exp:", c.Exp)
	s += fmt.Sprintf("%-8s %-10d %-8s %-8d\n", "level:", c.Level, "coin:", c.Coin)
	s += fmt.Sprintf("%-8s %-10s\n", "class:", c.Class)
	s += fmt.Sprintf("%-8s %-10s\n", "race:", c.Race)
	return s
}
func topRow(attr map[string]int, savings map[string]int) string {
	// colWdith := 16

	header := fmt.Sprintf("%-16s%-16s\n", "Attributes", "Savings")
	header += topBorder
	result := header
	result += fmt.Sprintf("%-4s%-12d%-14s%4d%8s\n", "str", attr["str"], "paralyze", savings["paralyze"], "|")
	result += fmt.Sprintf("%-4s%-12d%-14s%4d%8s\n", "int", attr["int"], "petrification", savings["petrification"], "|")
	result += fmt.Sprintf("%-4s%-12d%-14s%4d%8s\n", "wis", attr["wis"], "staff", savings["staff"], "|")
	result += fmt.Sprintf("%-4s%-12d%-14s%4d%8s\n", "dex", attr["dex"], "breath", savings["breath"], "|")
	result += fmt.Sprintf("%-4s%-12d%-14s%4d%8s\n", "con", attr["con"], "paralyze", savings["paralyze"], "|")
	result += fmt.Sprintf("%-4s%-12d%-14s%4d%8s\n", "cha", attr["cha"], "paralyze", savings["paralyze"], "|")
	return result
}
func (m *characterModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}
	return m, nil
}
