package tui

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	Space key.Binding
	Enter key.Binding
}

var keys = keyMap{
	Space: key.NewBinding(
		key.WithKeys("space"),
		key.WithHelp("␣", "alt screen"),
	),
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("↵", "open url"),
	),
}
