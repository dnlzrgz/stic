package hn

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"
)

func FetchItemsIds(c *http.Client, category string) ([]int, error) {
	url, err := categoryUrl(category)
	if err != nil {
		return nil, err
	}

	resp, err := c.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("got a %d status code", resp.StatusCode)
	}

	var ids []int
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, &ids); err != nil {
		return nil, err
	}

	return ids, nil
}

func FetchStory(c *http.Client, id int) (*Story, error) {
	jsonUrl, hnUrl := itemUrls(id)

	resp, err := c.Get(jsonUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var s Story
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, &s); err != nil {
		return nil, err
	}

	if s.Type == "story" && s.URL == "" {
		s.URL = hnUrl
	}

	return &s, nil
}

func FetchStories(c *http.Client, ids []int) (Stories, error) {
	if len(ids) <= 0 {
		return nil, errors.New("number of ids cannot be 0")
	}

	var wg sync.WaitGroup
	wg.Add(len(ids))

	stories := make(Stories, 0, len(ids))
	storiesCh := make(chan *Story, 5)
	for _, id := range ids {
		go func(id int) {
			defer wg.Done()

			s, _ := FetchStory(c, id) // TODO: handle error
			storiesCh <- s
		}(id)
	}

	go func() {
		wg.Wait()
		close(storiesCh)
	}()

	for s := range storiesCh {
		stories = append(stories, s)
	}

	return stories, nil
}
