package query_input

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/whoop-t/manly/internal/colors"
)

type Model struct {
	Input textinput.Model

	// If input is focused
	Focused bool
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func New() Model {
	input := textinput.New()
	input.Focus()
	input.Placeholder = "e.g bash"
	input.CharLimit = 50
	return Model{
		Input:   input,
		Focused: true, // Input is always focused first
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.Focused {
			m.Input, cmd = m.Input.Update(msg)
		}
	}
	return m, cmd
}

func (m Model) View(windowHeight int, windowWidth int) string {
	// Width calculated
	width := windowWidth / 3
	// Configure lipgloss styles for focused and unfocused states
	focusedStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colors.Purple)).
		Width(width).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#888888"))

	unfocusedStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#888888")).
		Width(width).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#888888"))

	// Determine which style to use based on focus
	var inputStyle lipgloss.Style
	if m.Focused {
		inputStyle = focusedStyle
	} else {
		inputStyle = unfocusedStyle
	}

	// Render the input view
	return inputStyle.Render(m.Input.View())
}
