package ui

type EndpointListLoadedMsg struct{}
type EndpointItemMsg struct {
	Message Item
}
type errorMsg struct {
	Message string
}

func errorCmd(msg string) tea.Cmd {
	return func() tea.Msg {
		return errorMsg{Message: msg}
	}
}
