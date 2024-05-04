package main

import (
	"strings"

	bcursor "github.com/charmbracelet/bubbles/cursor"
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
	canvas  string

	cursor  bcursor.Model
	cursorX int
	cursorY int
}

func newModel() model {
	cursor := bcursor.New()
	cursor.SetMode(bcursor.CursorStatic)
	return model{
		loading: true,
		cursor:  cursor,
	}
}

func (m *model) resize(w, h int) {
	m.width = w
	m.height = h

	ch := h - 2
	canvas := make([]string, ch)
	for i := 0; i < ch; i++ {
		canvas[i] = ""
	}
	m.canvas = strings.Join(canvas, "\n")

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
		}
	}
	m.cursor.Update(msg)
	return m, nil
}

func (m model) View() string {
	if m.loading {
		return "loading..."
	}
	status := lightStatus.Render("[status]")
	cmdline := "[cmd]"
	return lipgloss.JoinVertical(lipgloss.Top, m.canvas, status, cmdline)
}
