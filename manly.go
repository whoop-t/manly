package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

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
	list    list.Model
	input   textinput.Model
	focused int
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return bindings(m, msg)
}

func (m model) View() string {
	return fmt.Sprintf(
		"%s%s",
		m.input.View(),
		docStyle.Render(m.list.View()),
	)
}

func main() {
	ti := textinput.New()
	ti.Focus()
	ti.Placeholder = "Pikachu"
	ti.CharLimit = 156
	ti.Width = 20
	items := queryManPages("man")
	m := model{
		list:    list.New(items, list.NewDefaultDelegate(), 0, 0),
		input:   ti,
		focused: INPUT,
	}
	m.list.Title = "Man pages"

	p := tea.NewProgram(m)

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
