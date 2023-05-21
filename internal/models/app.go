package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/whoop-t/manly/internal/models/query_input"
)

/* App model is our main model state for the application
 * All other models will be constructed and returned to our App model
 */

type model struct {
	input query_input.Model
}

func New() model {
	return model {
		input: query_input.New(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// global bindings
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	// Render the view
	return m.input.View()
}
