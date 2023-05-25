package pager

import (
	"log"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/whoop-t/manly/internal/apis"
	"github.com/whoop-t/manly/internal/colors"
)

type Model struct {
	Page         viewport.Model
	IsPageSet    bool
	windowWidth  int
	windowHeight int
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.IsPageSet {
			if msg.String() == "j" {
				m.Page.LineDown(2)
			}
			if msg.String() == "k" {
				m.Page.LineUp(2)
			}
			if msg.String() == "q" || msg.String() == "esc" {
				m.IsPageSet = false
			}
		}
	case apis.ShowPageMessage:
		m.Page = NewPage(msg.Result, m.windowHeight, m.windowWidth)
		m.IsPageSet = true
		m.Page, cmd = m.Page.Update(msg)

	case tea.WindowSizeMsg:
		// Store window sizing
		m.windowWidth = msg.Width
		m.windowHeight = msg.Height
		m.Page.Style.Height(m.windowHeight).Width(m.windowWidth)
	}
	return m, cmd
}

func NewPage(content string, height int, width int) viewport.Model {
	page := viewport.New(width, height)
	page.Style = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(colors.Purple)).
		PaddingRight(2)

	renderer, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(width),
	)
	if err != nil {
		log.Fatal(err)
	}

	str, err := renderer.Render(content)
	if err != nil {
		log.Fatal(err)
	}

	page.SetContent(str)

	return page
}
