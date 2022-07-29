package main

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

var (
	primary    = lipgloss.Color("#ff6600")
	whiteColor = lipgloss.Color("#fffcf9")
	blackColor = lipgloss.Color("#000")
	helpColor  = lipgloss.Color("#f564a9")
)

var (
	baseStyle       = lipgloss.NewStyle().Padding(2).Foreground(whiteColor)
	titleStyle      = lipgloss.NewStyle().PaddingLeft(2).PaddingRight(4).Background(primary).Foreground(whiteColor)
	itemStyleDark   = lipgloss.NewStyle().PaddingLeft(2).Foreground(whiteColor)
	itemStyleLight  = itemStyleDark.Copy().Foreground(blackColor)
	paginationStyle = list.DefaultStyles().PaginationStyle.Foreground(helpColor)
	helpStyle       = list.DefaultStyles().HelpStyle.PaddingLeft(2).Foreground(helpColor)
)
