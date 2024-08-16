package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

type menuModel struct {
	choices  []character
	cursor   int
	selected map[int]struct{}
	width    int
	height   int
}

func initialModel(clist []character) menuModel {
	var choices []string
	for _, char := range clist {
		choices = append(choices, char.Name)
	}
	return menuModel{
		choices:  clist,
		selected: make(map[int]struct{}),
	}
}

func (m menuModel) Init() tea.Cmd {
	return nil
}

func (m menuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
                nextModel := NewApp(&m, m.width, m.height, m.choices[m.cursor])
				return &nextModel, nil
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}
	return m, nil
}

func (m menuModel) View() string {
	s := "choose a character:\n"
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice.Name)
	}
	s += "\nPress q to quit.\n"
	return s
}
