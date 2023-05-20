package main

import (
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

const debounceTimeSeconds = time.Second

type debounceMsg bool

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
		if msg.String() == "ctrl+j" || msg.String() == "down" {
			m.list.CursorDown()
		}
		if msg.String() == "ctrl+k" || msg.String() == "up" {
			m.list.CursorUp()
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
				m.list.ResetSelected()
			}
		}
		m.input, cmd = m.input.Update(msg)
	} else if m.focused == LIST {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			if msg.String() == "enter" {
			  // Need to get the actually page
			  
				page := m.list.SelectedItem().FilterValue()
				content := querySpecificPage(page)
			  m.page = newPage(content)
			  m.focused = PAGE
			}
		}
		m.list, cmd = m.list.Update(msg)
	} else if m.focused == PAGE {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			if msg.String() == "esc" || msg.String() == "q" {
			  m.focused = LIST
			}
		}
		m.page, cmd = m.page.Update(msg)
	}

	return m, cmd
}
