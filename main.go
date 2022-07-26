package main

import (
	"encoding/json"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/urfave/cli/v2"
	"log"
	"net/http"
	"os"
	"sort"
)

const hnUrl = "https://hacker-news.firebaseio.com"

var categories = map[string]string{
	"top":  "/v0/topstories.json",
	"news": "/v0/newstories.json",
	"best": "/v0/beststories.json",
	"ask":  "/v0/askstories.json",
	"show": "/v0/showstories.json",
	"job":  "/v0/jobstories.json",
}

func main() {
	var category string
	var maxItems int
	var debug bool
	var outputJson bool

	app := &cli.App{
		Name:                 "stic",
		Usage:                "hn in the terminal",
		EnableBashCompletion: true,
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
		},
		Action: func(ctx *cli.Context) error {
			path, ok := categories[category]
			if !ok {
				log.Fatalln("the category selected does not exists")
			}

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

			ids, err := fetchItemsIds(c, hnUrl+path)
			if err != nil {
				log.Fatalln(err)
			}

			if len(ids) < maxItems {
				maxItems = len(ids)
			} else {
				ids = ids[:maxItems]
			}

			stories, err := fetchStories(c, ids)
			if err != nil {
				log.Fatalln(err)
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
			m := newModel(category).withList(stories)
			if err := tea.NewProgram(m).Start(); err != nil {
				return err
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}
