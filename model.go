package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var lightStatus = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("0")).
	Background(lipgloss.Color("15"))

type model struct {
	loading bool
	width   int
	height  int

	controller controller
	canvas     canvas
}

func newModel() model {
	m := model{
		loading: true,
	}
	m.controller = newController()
	return m
}

func (m *model) resize(w, h int) {
	m.width = w
	m.height = h

	m.canvas = newCanvas(w, h)

	lightStatus.Width(w)
}

func (m *model) MoveCursor(count int, dir direction) {
	m.canvas.MoveCursor(count, dir)
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.loading = false
		m.resize(msg.Width, msg.Height)

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		default:
			m.controller.Handle(msg, &m)
		}
	}
	m.canvas.Update(msg)
	return m, nil
}

func (m model) View() string {
	if m.loading {
		return "loading..."
	}
	status := lightStatus.Render("[status]")
	cmdline := "[cmd]"
	canvas := m.canvas
	return lipgloss.JoinVertical(lipgloss.Top, canvas.asString(), status, cmdline)
}
