package app

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/whoop-t/manly/internal/apis"
	"github.com/whoop-t/manly/internal/models/pager"
	"github.com/whoop-t/manly/internal/models/query_input"
	"github.com/whoop-t/manly/internal/models/result_list"
	"github.com/whoop-t/manly/internal/models/title"
)

/* App model is our main model state for the application
 * All other models will be constructed and returned to our App model
 */

// Main application state model
type model struct {
	// Main query input for searching
	input query_input.Model

	// Main result with results from queries
	result result_list.Model

	// Pager currently showing for specific result
	pager      pager.Model
	isPageSet bool

	// Api that is being used, defaults to Man
	api apis.Api

	// Track height and width of the window for re-sizing
	windowHeight int
	windowWidth  int
}

// App state model constructor
func New() model {
	return model{
		input:  query_input.New(),
		result: result_list.New(),
		api:    apis.ManApi{},
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
		if msg.String() == "tab" {
			m.input.Focused = !m.input.Focused
			m.result.Focused = !m.result.Focused
		}
		if msg.String() == "enter" {
			if m.result.Focused {
				page := m.result.List.SelectedItem().FilterValue()
				sCmd := m.api.ShowPage(page)
				return m, sCmd
			} else {
				glCmd := m.api.GetList(m.input.Input.Value())
				return m, glCmd
			}
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
	var c tea.Cmd
	m.input, a = m.input.Update(msg)
	m.result, b = m.result.Update(msg)
	m.pager, b = m.pager.Update(msg)
	cmds = append(cmds, a)
	cmds = append(cmds, b)
	cmds = append(cmds, c)
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	var view string
	if m.pager.IsPageSet {
		view = m.pager.Page.View()
	} else {
		// Render the view, cascade to all children View methods
		view = fmt.Sprintf(
			"%s\n\n%s\n\n%s",
			title.View(),
			m.input.View(),
			m.result.View(),
		)
	}

	return view
}
