package main

import (
	"strconv"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

// Definitions

type direction int

const (
	DIRECTION_UP direction = iota
	DIRECTION_DOWN
	DIRECTION_LEFT
	DIRECTION_RIGHT
)

type listener interface {
	MoveCursor(count int, dir direction)
}

type mode interface {
	Handle(msg tea.KeyMsg, listener listener)
}

// KeyMaps for different modes

type normalModeKeyMap struct {
	MoveUp    key.Binding
	MoveDown  key.Binding
	MoveLeft  key.Binding
	MoveRight key.Binding

	Digits key.Binding
}

// Modes

type normalMode struct {
	repCount int
	keyMap   normalModeKeyMap
}

func newNormalMode() normalMode {
	return normalMode{
		keyMap: normalModeKeyMap{

			MoveUp: key.NewBinding(
				key.WithKeys("up", "k"),
				key.WithHelp("↑/k", "Up"),
			),
			MoveDown: key.NewBinding(
				key.WithKeys("down", "j"),
				key.WithHelp("↓/j", "Down"),
			),
			MoveLeft: key.NewBinding(
				key.WithKeys("left", "h"),
				key.WithHelp("←/h", "Left"),
			),
			MoveRight: key.NewBinding(
				key.WithKeys("right", "l"),
				key.WithHelp("→/l", "Right"),
			),

			Digits: key.NewBinding(
				key.WithKeys("0", "1", "2", "3", "4", "5", "6", "7", "8", "9"),
			),
		},
	}
}

func (m *normalMode) Handle(msg tea.KeyMsg, listener listener) {
	var count int
	if m.repCount == 0 {
		count = 1
	} else {
		count = m.repCount
	}
	clearCount := true

	switch {

	case key.Matches(msg, m.keyMap.MoveUp):
		listener.MoveCursor(count, DIRECTION_UP)
	case key.Matches(msg, m.keyMap.MoveDown):
		listener.MoveCursor(count, DIRECTION_DOWN)
	case key.Matches(msg, m.keyMap.MoveLeft):
		listener.MoveCursor(count, DIRECTION_LEFT)
	case key.Matches(msg, m.keyMap.MoveRight):
		listener.MoveCursor(count, DIRECTION_RIGHT)

	case key.Matches(msg, m.keyMap.Digits):
		clearCount = false
		digit, err := strconv.Atoi(msg.String())
		if err == nil {
			m.repCount = m.repCount*10 + digit
		}
	}

	if clearCount {
		m.repCount = 0
	}
}

// Controller

type controller struct {
	currentMode mode

	normalMode normalMode
}

func newController() controller {
	normalMode := newNormalMode()
	return controller{
		currentMode: &normalMode,
		normalMode:  normalMode,
	}
}

func (c *controller) Handle(msg tea.KeyMsg, listener listener) {
	c.currentMode.Handle(msg, listener)
}
