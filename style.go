package main

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

const (
	orangeColor = "#ff6600"
	whiteColor  = "#fff"
	textColor   = "#000"
	helpColor   = "#828282"
)

var (
	baseStyle       = lipgloss.NewStyle().Padding(2).Foreground(lipgloss.AdaptiveColor{Light: textColor, Dark: whiteColor})
	titleStyle      = lipgloss.NewStyle().PaddingLeft(2).PaddingRight(4).Background(lipgloss.Color(orangeColor)).Foreground(lipgloss.Color(whiteColor))
	paginationStyle = list.DefaultStyles().PaginationStyle.Foreground(lipgloss.AdaptiveColor{Light: helpColor, Dark: whiteColor})
	helpStyle       = list.DefaultStyles().HelpStyle.PaddingLeft(2).MarginBottom(2).Foreground(lipgloss.AdaptiveColor{Light: helpColor, Dark: whiteColor})
)
