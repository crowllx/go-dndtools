package main

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type appModel struct {
	width     int
	height    int
	prevModel *menuModel
	character *characterModel
	session   *sessionModel
	ti        textinput.Model
}

func NewApp(prev *menuModel, width, height int, data character) appModel {
	app := appModel{
		width:     width,
		height:    height,
		prevModel: prev,
		character: &characterModel{
			data:   data,
			width:  width,
			height: height,
		},
		session: nil,
		ti:      textinput.New(),
	}
	app.ti.Cursor.Style.Blink(true)
	app.ti.CharLimit = 50
	app.ti.Width = width
	app.ti.Cursor.Blink = true
	app.ti.Focus()
	return app
}

func (m *appModel) Init() tea.Cmd {
	return tea.SetWindowTitle("dnd tools!")

}
func (m *appModel) View() string {
	var s string
	if m.session != nil {
		s += lipgloss.JoinHorizontal(lipgloss.Top, m.character.View(), m.session.View())
	} else {
		s += m.character.View()
	}
	s += "\n" + m.ti.View()
	return s
}
func (m *appModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m.prevModel, nil
		case tea.KeyEnter:
			if m.session != nil {
				m.session.events = append(m.session.events, m.ti.Value())
			} else {
				m.session = &sessionModel{
					events: []string{
						m.ti.Value(),
					},
				}
			}
			m.ti.Reset()
		}
	}
	m.ti, cmd = m.ti.Update(msg)
	model, _ := m.character.Update(msg)
	m.character = model.(*characterModel)
	return m, cmd
}
