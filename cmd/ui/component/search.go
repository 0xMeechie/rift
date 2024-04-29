package component

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type SearchModel struct {
	textBox textinput.Model
}

var ()

func (s SearchModel) Init() tea.Cmd {
	return textinput.Blink
}
func (s SearchModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "/":
			return s, s.textBox.Focus()
		}
	}

	return s, nil
}
func (s SearchModel) View() string {
	return helpStyle.Render(s.textBox.View())
}

func InitSearch() textinput.Model {
	ti := textinput.New()
	ti.Placeholder = "What are we listening to?"
	ti.CharLimit = 150

	return ti
}
