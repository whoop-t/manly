package query_input

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	input textinput.Model
}

func Init() tea.Msg {
  return textinput.Blink()
}

func New() Model {
	input := textinput.New()
	input.Focus()
	input.Placeholder = "e.g bash"
	input.CharLimit = 50
	return Model{
	  input: input,
	}
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
