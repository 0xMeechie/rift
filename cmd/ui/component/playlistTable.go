package component

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fdaygon/rift/pkg/spotify"
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
		switch {
		//	case key.Matches(msg, DefaultKeys.ESC):
		//		if p.Table.Focused() {
		//			p.Table.Blur()
		//		} else {
		//			p.Table.Focus()
		//		}

		case key.Matches(msg, DefaultKeys.Up):
			fmt.Println("up")
			p.Table.MoveUp(1)
		case key.Matches(msg, DefaultKeys.Down):
			p.Table.MoveDown(1)

		}
	}
	return p, cmd
}

func InitPlaylist() table.Model {
	//Create the columns to the playlist
	colums := []table.Column{
		{Title: "#", Width: 5},
		{Title: "Playlist", Width: 20},
		{Title: "Total Songs", Width: 15},
	}
	rows := []table.Row{}
	//Get the users playlist and add them to the table
	playlists := spotify.GetPlaylist()

	for idx, playlist := range playlists.Items {
		row := []string{
			fmt.Sprint(idx + 1),
			playlist.Name,
			fmt.Sprint(playlist.Tracks.Total),
			playlist.ID,
		}
		rows = append(rows, row)
	}

	modelTable := table.New(
		table.WithColumns(colums),
		table.WithRows(rows),
		table.WithWidth(50),
		table.WithFocused(true),
	)
	style := table.DefaultStyles()

	modelTable.SetStyles(style)

	return modelTable

}
