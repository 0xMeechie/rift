package ui

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
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
	Table       table.Model
	SongTable   table.Model
	Help        component.HelpModel
	Search      textinput.Model
}

func (m homeModel) Init() tea.Cmd {
	return nil
}
func (m homeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "/":
			if m.SessionView != searchView {
				m.SessionView = searchView
				m.Search, cmd = m.Search.Update(msg)

			}
		case "tab":
			if m.SessionView == playlistView {

				m.SessionView = songView
			} else if m.SessionView == songView {
				m.SessionView = playlistView
			}
		case "esc":
			if m.SessionView == playlistView {
				if m.Table.Focused() {
					m.Table.Blur()
				} else {
					m.Table.Focus()
				}

			} else if m.SessionView == songView {
				if m.SongTable.Focused() {
					m.SongTable.Blur()
				} else {
					m.SongTable.Focus()
				}
			} else if m.SessionView == searchView {
				m.SessionView = playlistView
			}
		case "enter":
			if m.SessionView == playlistView {
				selectedPlaylist := m.Table.SelectedRow()
				fmt.Println(selectedPlaylist[0])
			}

		}

		switch m.SessionView {
		case playlistView:

			m.Table, cmd = m.Table.Update(msg)
			//		case songView:
			//			m.SongTable, cmd = m.SongTable.Update(msg)
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
