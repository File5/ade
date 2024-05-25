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

	canvas canvas
}

func newModel() model {
	return model{
		loading: true,
	}
}

func (m *model) resize(w, h int) {
	m.width = w
	m.height = h

	m.canvas = newCanvas(w, h)

	lightStatus.Width(w)
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
		case "t":
			m.canvas.setCursorPos(12, 3)
		case "up", "k":
			m.canvas.MoveCursor(1, CANVAS_DIR_UP)
		case "down", "j":
			m.canvas.MoveCursor(1, CANVAS_DIR_DOWN)
		case "left", "h":
			m.canvas.MoveCursor(1, CANVAS_DIR_LEFT)
		case "right", "l":
			m.canvas.MoveCursor(1, CANVAS_DIR_RIGHT)
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
