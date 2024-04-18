package component

import "github.com/charmbracelet/bubbles/key"

type keymap struct {
	Up     key.Binding
	Down   key.Binding
	Left   key.Binding
	Right  key.Binding
	Help   key.Binding
	Quit   key.Binding
	Search key.Binding
}

func (k keymap) ShortHelp() []key.Binding {
	return nil
}

func (k keymap) FullHelp() [][]key.Binding {
	return nil
}

var defaultKeys = keymap{
	Up: key.NewBinding(
		key.WithKeys("k", "up"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.Binding,
}
