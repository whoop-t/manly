package main

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func bindings(m model, msg tea.Msg) (model, tea.Cmd) {
	// global bindings
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		if msg.String() == "tab" {
			if m.focused == LIST {
				m.focused = INPUT
			} else {
				m.focused = LIST
			}
			return m, textinput.Blink
		}
		if msg.String() == "ctrl+j" {
			m.list.CursorDown()
		}
		if msg.String() == "ctrl+k" {
			m.list.CursorUp()
		}
		if msg.String() == "ctrl+l" {
			m.list.Paginator.NextPage()
		}
		if msg.String() == "ctrl+h" {
			m.list.Paginator.PrevPage()
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	if m.focused == INPUT {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			if msg.String() == "enter" {
				m.list.SetItems(queryManPages(m.input.Value()))
			}
		}
		m.input, cmd = m.input.Update(msg)
	} else if m.focused == LIST {
		m.list, cmd = m.list.Update(msg)
	}

	return m, cmd
}
