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
	list list.Model
}

func New() Model {
	m := Model{
		list: list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0), // Empty list with defaults
	}
	m.list.SetHeight(20)
	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "enter" {
		}
	case apis.ListFetchedMessage:
		cmd = m.list.SetItems(formatList(msg.Results))
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
	return m.list.View()
}
