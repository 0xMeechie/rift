package component

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

type SongTable struct {
	Table table.Model
}

func (s SongTable) Init() tea.Cmd {
	return nil
}
func (s SongTable) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return s, tea.Quit
}
func (s SongTable) View() string {
	return s.Table.View()
}

func InitSong() table.Model {
	columns := []table.Column{
		{Title: "#", Width: 5},
		{Title: "Title", Width: 25},
		{Title: "Album", Width: 15},
		{Title: "Date Added", Width: 10},
		{Title: "Length", Width: 8},
	}
	rows := []table.Row{
		{"1", "Slow Dancing in the Dark", "Album 1", "April 17", "3:27"},
		{"2", "Die for you", "Beauty behind the madness ", "May 6", "4:27"},
		{"3", "Take Care", "Album 3", "August 1", "3:24"},
		{"4", "Slow Dancing in the Dark", "Album 5", "September 16", "2:13"},
		{"5", "Slow Dancing in the Dark", "Album 8", "March 1", "4:03"},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithHeight(20),
		table.WithFocused(true),
	)
	return t
}
