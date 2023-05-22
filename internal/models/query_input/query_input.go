package query_input

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/whoop-t/manly/internal/apis"
)

type Model struct {
	input textinput.Model
	api   apis.Api
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func New(api apis.Api) Model {
	input := textinput.New()
	input.Focus()
	input.Placeholder = "e.g bash"
	input.CharLimit = 50
	return Model{
		input: input,
		api: api,
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "enter" {
			te := m.api.GetList(m.input.Value())
			return m, te
		}
	}
	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	// Configure lipgloss styles for focused and unfocused states
	focusedStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")).
		Width(30).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#888888"))

	// unfocusedStyle := lipgloss.NewStyle().
	// 	Foreground(lipgloss.Color("#888888")).
	// 	Width(30).
	// 	Border(lipgloss.RoundedBorder()).
	// 	BorderForeground(lipgloss.Color("#888888"))

	// Determine which style to use based on focus
	var inputStyle lipgloss.Style
	// if m.focused == INPUT {
	// 	inputStyle = focusedStyle
	// } else {
	// 	inputStyle = unfocusedStyle
	// }

	inputStyle = focusedStyle
	// Render the input view
	return inputStyle.Render(m.input.View())
}
