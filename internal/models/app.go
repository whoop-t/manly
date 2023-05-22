package app

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/whoop-t/manly/internal/apis"
	"github.com/whoop-t/manly/internal/models/query_input"
	"github.com/whoop-t/manly/internal/models/result_list"
)

/* App model is our main model state for the application
 * All other models will be constructed and returned to our App model
 */

// Main application state model
type model struct {
	// Main query input for searching
	input query_input.Model

	// Main list with results from queries
	list result_list.Model

	// Track height and width of the window for re-sizing
	windowHeight int
	windowWidth  int
}

// App state model constructor
func New() model {
	return model{
		input: query_input.New(apis.ManApi{}),
		list:  result_list.New(),
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

	case tea.WindowSizeMsg:
		// Store window sizing
		m.windowWidth = msg.Width
		m.windowHeight = msg.Height
	}

	/* TODO theres gotta be a better way to batch those*/
	/* or maybe batch isnt even needed...*/
	var cmds []tea.Cmd
	var a tea.Cmd
	var b tea.Cmd

	// TODO need to cascade updates to all components
	m.input, a = m.input.Update(msg)
	m.list, b = m.list.Update(msg)
	cmds = append(cmds, a)
	cmds = append(cmds, b)
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	// Render the view, cascade to all children View methods
	view := fmt.Sprintf("%s%s", m.input.View(), m.list.View())
	return view
}
