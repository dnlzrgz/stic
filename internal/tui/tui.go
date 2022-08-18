package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/daniarlert/stic/internal/hn"
	"github.com/pkg/browser"
)

type model struct {
	category   string
	totalItems int
	altScreen  bool
	spinner    spinner.Model
	list       list.Model
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
		case tea.KeyCtrlC, tea.KeyEscape:
			return m, tea.Quit
		case tea.KeyEnter:
			s, ok := m.list.SelectedItem().(item)
			if !ok {
				return m, tea.Quit
			}

			// TODO: send tea.Msg with an error
			if len(s.desc) == 0 {
				return m, cmd
			}

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

func NewModel(category string, max int) model {
	if _, ok := hn.Categories[category]; !ok {
		// TODO: don't panic
		panic("the selected category does not exists")
	}

	m := model{
		category:   category,
		totalItems: max,
	}

	return m
}

func (m model) WithSpinner() model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = spinnerStyle

	m.spinner = s
	return m
}

func (m model) WithList(stories hn.Stories) model {
	items := make([]list.Item, 0)
	for _, s := range stories {
		items = append(items, item{
			title: fmt.Sprintf("↑%d - %q by %s", s.Score, s.Title, s.By),
			desc:  s.URL,
		})
	}

	l := list.New(items, itemDelegate{itemStyleDark}, 0, 0)
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

func (m model) WithLightColors() model {
	m.list.SetDelegate(itemDelegate{itemStyleLight})
	return m
}
