package apis

import tea "github.com/charmbracelet/bubbletea"

// Api contract
// We want all apis to adhere to this interface
type Api interface {
	GetList(query string) tea.Cmd
	ShowPage(query string) tea.Cmd
}

// ListFetchedMessage is a message when the api completes the query
type ListFetchedMessage struct {
	Results []string
}

// ShowPageMessage is a message when the api completes the query
type ShowPageMessage struct {
	Result string
}
