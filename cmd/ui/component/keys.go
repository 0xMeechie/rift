package component

import (
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

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
	return []key.Binding{k.Quit, k.Help}
}

func (k keymap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up},
		{k.Down},
		{k.Left},
		{k.Right},
		{k.Search},
		{k.Help},
		{k.Quit},
	}
}

var defaultKeys = keymap{
	Up: key.NewBinding(
		key.WithKeys("k", "up"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("j", "down"),
		key.WithHelp("↓/j", "move down"),
	),
	Left: key.NewBinding(
		key.WithKeys("h", "left"),
		key.WithHelp("/h", "move left"),
	),
	Right: key.NewBinding(
		key.WithKeys("l", "right"),
		key.WithHelp("/l", "move right"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "Get help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl-c"),
		key.WithHelp("q/esc/ctrl-c", "quit"),
	),
	Search: key.NewBinding(
		key.WithKeys("/"),
		key.WithHelp("/", "Search"),
	),
}

type HelpModel struct {
	key  keymap
	help help.Model
}

func NewHelpModel() HelpModel {
	return HelpModel{
		key:  defaultKeys,
		help: help.New(),
	}
}

func (h HelpModel) Init() tea.Cmd {
	return nil
}
func (h HelpModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return h, tea.Quit
}
func (h HelpModel) View() string {

	h.help.ShowAll = true
	helpView := h.help.View(h.key)
	height := strings.Count(helpView, "\n")
	return "\n" + strings.Repeat("\n", height) + helpView
}
