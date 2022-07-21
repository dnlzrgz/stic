package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"net/http"
	"os"
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

	app := &cli.App{
		Name:                 "stic",
		Usage:                "navigate HN in the terminal",
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "category",
				Aliases:     []string{"c"},
				Value:       "top",
				Usage:       "HN category",
				Destination: &category,
			},
			&cli.IntFlag{
				Name:        "max",
				Aliases:     []string{"m"},
				Value:       0,
				Usage:       "max number of items",
				Destination: &maxItems,
			},
		},
		Action: func(ctx *cli.Context) error {
			if _, ok := categories[category]; !ok {
				log.Fatalln("the category selected does not exists")
			}

			path := categories[category]
			log.Println("category " + category + " is on path " + path)

			c := &http.Client{}

			ids, err := fetchItemsIds(c, hnUrl+path)
			if err != nil {
				log.Fatalln(err)
			}

			for i, id := range ids {
				fmt.Println(i+1, id)
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}
