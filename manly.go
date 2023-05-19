package main

import (
	"log"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

type (
	errMsg error
)

type model struct {
	searchInput textinput.Model
	resultsList list.Model
	err         error
}

func initialModel() model {
	searchInput := textinput.New()
	searchInput.Placeholder = "which"
	searchInput.Focus()
	searchInput.CharLimit = 40
	searchInput.Width = 20
	
	// Init model
	m := model{
		resultsList:  list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0),
		searchInput: searchInput,
		err:         nil,
	}
	
	m.resultsList.SetStatusBarItemName("man", "mans")
	//Hide default list title + styles
	m.resultsList.SetShowTitle(false)
	return m
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			m.resultsList.SetItems(queryManPages(m.searchInput.Value()))
			return m, nil
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.searchInput, cmd = m.searchInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return render(m)
}
