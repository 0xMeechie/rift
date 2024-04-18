package ui

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type homeModel struct {
	Table table.Model
}

func (m homeModel) Init() tea.Cmd {
	return nil
}
func (m homeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		}
	}
	return m, cmd
}

func (m homeModel) View() string {
	return m.Table.View()
}
func InitModel() {
	colums := []table.Column{
		{Title: "Playlist", Width: 10},
		{Title: "Total Songs", Width: 40},
	}

	rows := []table.Row{
		{"Chill Coding Playlist", "312"},
		{"Old Vibes", "123"},
		{"Liked Songs", "4532"},
	}

	modelTable := table.New(
		table.WithColumns(colums),
		table.WithRows(rows),
		table.WithWidth(1080),
	)

	style := table.DefaultStyles()

	style.Header = style.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)

	style.Selected = style.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)

	modelTable.SetStyles(style)

	m := homeModel{modelTable}

	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
