package ui

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/fdaygon/rift/cmd/ui/component"
)

var (
	tableStyle = lipgloss.NewStyle().BorderStyle(lipgloss.DoubleBorder())
)

type sessionState uint

const (
	playlistView sessionState = iota + 1
	songView
	searchView
)

type homeModel struct {
	SessionView sessionState
	Table       tea.Model
	SongTable   tea.Model
	Help        component.HelpModel
	Search      tea.Model
}

func (m homeModel) Init() tea.Cmd {
	return nil
}
func (m homeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	fmt.Println(msg)
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		case "/":
			if m.SessionView != searchView {
				m.SessionView = searchView
				m.Search, cmd = m.Search.Update(msg)

			}
		}

		switch m.SessionView {
		case playlistView:
			m.Table, cmd = m.Table.Update(msg)
		case songView:
			m.SongTable, cmd = m.SongTable.Update(msg)
		}
	}
	return m, cmd
}

func (m homeModel) View() string {

	switch m.SessionView {

	case playlistView:

		return lipgloss.PlaceHorizontal(100, 25, m.Help.View()) + "\n\n" + lipgloss.PlaceHorizontal(100, 20, m.Search.View()) + "\n\n" + lipgloss.JoinHorizontal(0.2, tableStyle.Render(m.Table.View()), tableStyle.Render(m.SongTable.View()))
	case songView:
		return "Song View"
	case searchView:
		return "Search View"
	}

	return "Error Getting View"

}

func InitModel() {
	colums := []table.Column{
		{Title: "Playlist", Width: 20},
		{Title: "Total Songs", Width: 15},
	}

	rows := []table.Row{
		{"Chill Coding Playlist", "312"},
		{"Old Vibes", "123"},
		{"Liked Songs", "4532"},
	}

	modelTable := table.New(
		table.WithColumns(colums),
		table.WithRows(rows),
		table.WithWidth(50),
	)
	style := table.DefaultStyles()

	style.Header = tableStyle

	modelTable.SetStyles(style)

	m := homeModel{
		SessionView: playlistView,
		Table:       component.InitPlaylist(),
		Help:        component.NewHelpModel(),
		Search:      component.InitSearch(),
		SongTable:   component.InitSong(),
	}

	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
