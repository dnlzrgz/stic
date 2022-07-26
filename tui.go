package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/pkg/browser"
	"io"
)

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

type item struct {
	title string
	desc  string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, _ list.Model, _ int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%s\n%s↵", i.title, i.desc)

	fn := selectedItemStyle.Render
	fmt.Fprintf(w, fn(str))
}

type model struct {
	category  string
	altScreen bool
	list      list.Model
	items     []item
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyEnter:
			s, ok := m.list.SelectedItem().(item)
			if !ok {
				return m, tea.Quit
			}

			// TODO: handle error
			if err := browser.OpenURL(s.desc); err != nil {
				return m, tea.Quit
			}
		case tea.KeySpace:
			if m.altScreen {
				cmd = tea.ExitAltScreen
			} else {
				cmd = tea.EnterAltScreen
			}

			m.altScreen = !m.altScreen
			return m, cmd
		}
	}

	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return baseStyle.Render(m.list.View() + "\n")
}

func newModel(category string) model {
	m := model{
		category: category,
	}

	return m
}

func (m model) withList(stories stories) model {
	items := make([]list.Item, 0, len(stories))
	for _, s := range stories {
		items = append(items, item{
			title: fmt.Sprintf("↑%d - %q by %s", s.Score, s.Title, s.By),
			desc:  s.URL,
		})
	}

	l := list.New(items, itemDelegate{}, 0, 0)
	l.Title = fmt.Sprintf("hn - %s stories", m.category)

	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)

	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	l.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{keys.Space, keys.Enter}
	}

	l.AdditionalFullHelpKeys = func() []key.Binding {
		return []key.Binding{keys.Space, keys.Enter}
	}

	m.list = l

	return m
}
