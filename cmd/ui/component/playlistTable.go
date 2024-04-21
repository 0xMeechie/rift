package component

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

type PlaylistTable struct {
	Table table.Model
}

func (p PlaylistTable) Init() tea.Cmd {
	return nil
}

func (p PlaylistTable) View() string {
	return p.Table.View()
}

func (p PlaylistTable) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return p, tea.Quit
		}
	}
	return p, cmd
}

func InitPlaylist() PlaylistTable {
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

	modelTable.SetStyles(style)

	m := PlaylistTable{
		Table: modelTable,
	}
	return m

}
