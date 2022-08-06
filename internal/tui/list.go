package tui

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type item struct {
	title string
	desc  string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return "" }

type itemDelegate struct {
	itemStyle lipgloss.Style
}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, _ list.Model, _ int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	out := fmt.Sprintf("%s\n%s↲", i.title, i.desc)
	fmt.Fprint(w, d.itemStyle.Render(out))
}
