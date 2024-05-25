package main

import (
	"strings"

	bcursor "github.com/charmbracelet/bubbles/cursor"
	tea "github.com/charmbracelet/bubbletea"
)

type canvas struct {
	width  int
	height int

	content string

	cursor  bcursor.Model
	cursorX int
	cursorY int
}

func newCanvas(width, height int) canvas {
	cursor := bcursor.New()
	// cursor.TextStyle = lipgloss.NewStyle().
	// 	Foreground(lipgloss.Color("0")).
	// 	Background(lipgloss.Color("15"))
	cursor.SetMode(bcursor.CursorStatic)
	cursor.SetChar(" ")
	cursor.Focus()
	canvas := canvas{
		width:  width,
		height: height,
		cursor: cursor,
	}
	canvas.setCursorPos(0, 0)
	return canvas
}

func (c *canvas) setCursorPos(x, y int) {
	c.cursorX = x
	c.cursorY = y

	ch := c.height - 2
	canvas := make([]string, ch)
	for i := 0; i < ch; i++ {
		if i == c.cursorY {
			canvas[i] = strings.Repeat(" ", c.cursorX) +
				c.cursor.View()
		} else {
			canvas[i] = ""
		}
	}
	c.content = strings.Join(canvas, "\n")
}

func (c *canvas) asString() string {
	return c.content
}

func (c *canvas) Update(msg tea.Msg) (bcursor.Model, tea.Cmd) {
	return c.cursor.Update(msg)
}
