package result_list

import (
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/whoop-t/manly/internal/apis"
)

// Basic item struct for values in the list
type item struct {
	title string
	desc  string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type Model struct {
	List list.Model

	// If list is focused
	Focused bool
}

func New() Model {
	m := Model{
		List:    list.New([]list.Item{}, list.NewDefaultDelegate(), 20, 20), // Empty list with defaults
		Focused: false,
	}
	m.List.SetShowTitle(false)
	m.List.SetShowHelp(false)
	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.Focused {
			if msg.String() == "j" {
				m.List.CursorDown()
			}
			if msg.String() == "k" {
				m.List.CursorUp()
			}
			m.List, cmd = m.List.Update(cmd)
		}
	case apis.ListFetchedMessage:
		cmd = m.List.SetItems(formatList(msg.Results))
	}
	return m, cmd
}

// Format the strings to title/desc pairs
func formatList(results []string) []list.Item {
	// Convert list to []list.Item
	var list []list.Item
	for _, result := range results {
		// We also format the return man string so we can have title and desc
		parts := strings.Split(result, " - ")
		if len(parts) != 2 {
			continue
		}
		list = append(list, item{title: parts[0], desc: parts[1]})
	}
	return list
}

func (m Model) View() string {
	return m.List.View()
}
