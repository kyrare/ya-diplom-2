package bubbletea

import "github.com/charmbracelet/lipgloss"

const (
	hotPink  = lipgloss.Color("#EE6FF8")
	darkGray = lipgloss.Color("#767676")
	red      = lipgloss.Color("#EE204D")
	green    = lipgloss.Color("#5fb458")
)

var (
	docStyle   = lipgloss.NewStyle().Margin(1, 2)
	titleStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("62")).
			Foreground(lipgloss.Color("230")).
			Padding(0, 1)
	inputLabelStyle = lipgloss.NewStyle().Foreground(hotPink)
	errorStyle      = lipgloss.NewStyle().Foreground(red)
	successStyle    = lipgloss.NewStyle().Foreground(green)
	continueStyle   = lipgloss.NewStyle().Foreground(darkGray)
)
