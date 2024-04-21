package component

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type SearchModel struct {
	textBox textinput.Model
}

func (s SearchModel) Init() tea.Cmd {
	return textinput.Blink
}
func (s SearchModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return s, nil
}
func (s SearchModel) View() string {
	return s.textBox.View()
}

func InitSearch() SearchModel {
	ti := textinput.New()
	ti.Placeholder = "What are we listening to?"
	ti.CharLimit = 150
	model := SearchModel{
		textBox: ti,
	}
	return model
}
