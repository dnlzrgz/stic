package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/pkg/browser"
)

type keyMap struct {
	Space key.Binding
	Enter key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Space, k.Enter}
}

func (k keyMap) FullHelp() []key.Binding {
	return []key.Binding{k.Space, k.Enter}
}

var keys = keyMap{
	Space: key.NewBinding(
		key.WithKeys("space"),
		key.WithHelp("␣", "toggle alt screen"),
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

type model struct {
	category  string
	altScreen bool
	keys      keyMap
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
		keys:     keys,
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

	delegate := list.NewDefaultDelegate()
	delegate.Styles.SelectedTitle.Foreground(lipgloss.Color(textColor))

	l := list.New(items, delegate, 0, 0)
	l.Title = fmt.Sprintf("hn - %s stories", m.category)

	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)

	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	l.AdditionalShortHelpKeys = m.keys.ShortHelp
	l.AdditionalFullHelpKeys = m.keys.FullHelp

	m.list = l

	return m
}
