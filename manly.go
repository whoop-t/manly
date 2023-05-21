package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

const (
	INPUT = 0
	LIST  = 1
	PAGE  = 2
)

type model struct {
	list         list.Model
	input        textinput.Model
	page         viewport.Model
	focused      int
	windowWidth  int
	windowHeight int
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return bindings(m, msg)
}

func (m model) View() string {
	// Render the view
	return render(m)
}

func main() {
	ti := textinput.New()
	ti.Focus()
	ti.Placeholder = "e.g bash"
	ti.CharLimit = 50
	m := model{
		list:    list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0),
		input:   ti,
		page:    viewport.Model{}, // init empty page
		focused: INPUT,
	}
	m.list.SetShowTitle(false)

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
