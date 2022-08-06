package hn

import (
	"errors"
	"fmt"
	"net/url"
)

const baseUrl = "https://hacker-news.firebaseio.com"

var Categories = map[string]string{
	"top":  "/v0/topstories.json",
	"news": "/v0/newstories.json",
	"best": "/v0/beststories.json",
	"ask":  "/v0/askstories.json",
	"show": "/v0/showstories.json",
	"job":  "/v0/jobstories.json",
}

func categoryUrl(category string) (string, error) {
	path, ok := Categories[category]
	if !ok {
		return "", errors.New("received category does not exists")
	}

	return url.JoinPath(baseUrl, path)
}

func itemUrls(id int) (string, string) {
	return fmt.Sprintf("%s/v0/item/%v.json", baseUrl, id), fmt.Sprintf("https://news.ycombinator.com/item?id=%v", id)
}
