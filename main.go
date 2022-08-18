package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/daniarlert/stic/internal/hn"
	"github.com/daniarlert/stic/internal/tui"
	"github.com/urfave/cli/v2"
)

func main() {
	if err := start(os.Args); err != nil {
		log.Fatalf("error while running stic: %v", err)
	}
}

func start(args []string) error {
	var category string
	var maxItems int
	var noLimit bool
	var debug bool
	var outputJson bool
	var lightMode bool

	app := &cli.App{
		Name:                 "stic",
		Usage:                "hn in the terminal",
		EnableBashCompletion: true,
		Version:              "v0.0.4",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "category",
				Aliases:     []string{"c"},
				Value:       "top",
				Usage:       "hn category. Available categories are: \"top\", \"news\", \"best\", \"ask\", \"show\" and \"job\"",
				Destination: &category,
			},
			&cli.IntFlag{
				Name:        "max",
				Aliases:     []string{"m"},
				Value:       20,
				Usage:       "max number of items",
				Destination: &maxItems,
			},
			&cli.BoolFlag{
				Name:        "no-limit",
				Value:       false,
				Usage:       "fetches as many stories as possible ignoring the `--max` flag",
				Destination: &noLimit,
			},
			&cli.BoolFlag{
				Name:        "json",
				Value:       false,
				Usage:       "outputs JSON object",
				Destination: &outputJson,
			},
			&cli.BoolFlag{
				Name:        "debug",
				Aliases:     []string{"d"},
				Value:       false,
				Usage:       "enables debug mode",
				Destination: &debug,
			},
			&cli.BoolFlag{
				Name:        "light",
				Value:       false,
				Usage:       "enables light color scheme",
				Destination: &lightMode,
			},
		},
		Action: func(ctx *cli.Context) error {
			if debug {
				f, err := tea.LogToFile("stic.log", "debug")
				if err != nil {
					log.Fatalln(err)
				} else {
					defer func() {
						if err := f.Close(); err != nil {
							log.Fatalln(err)
						}
					}()
				}
			}

			c := &http.Client{}

			ids, err := hn.FetchItemsIds(c, category)
			if err != nil {
				return err
			}

			if !noLimit {
				if len(ids) < maxItems {
					maxItems = len(ids)
				} else {
					ids = ids[:maxItems]
				}
			}

			stories, err := hn.FetchStories(c, ids)
			if err != nil {
				return err
			}

			if outputJson {
				jsonObj, err := json.Marshal(stories)
				if err != nil {
					return err
				}

				fmt.Fprintf(os.Stdout, "%s", string(jsonObj))
				return nil
			}

			sort.Sort(stories)
			m := tui.NewModel(category, maxItems).WithSpinner().WithList(stories)

			if lightMode {
				m = m.WithLightColors()
			}

			if err := tea.NewProgram(m).Start(); err != nil {
				return err
			}

			return nil
		},
	}

	return app.Run(args)
}
