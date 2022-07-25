package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func fetchItemsIds(c *http.Client, url string) ([]int, error) {
	resp, err := c.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error request responde with status code %d", resp.StatusCode)
	}

	var ids []int
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, &ids); err != nil {
		return nil, err
	}

	return ids, nil
}

func fetchStory(c *http.Client, url string) (*story, error) {
	resp, err := c.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var s story
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, &s); err != nil {
		return nil, err
	}

	if s.Type == "story" && s.URL == "" {
		s.URL = fmt.Sprintf("https://news.ycombinator.com/item?id=%v", s.ID)
	}

	return &s, nil
}

func fetchStories(c *http.Client, ids []int) (stories, error) {
	if len(ids) <= 0 {
		return nil, errors.New("number of ids cannot be 0")
	}

	var wg sync.WaitGroup
	wg.Add(len(ids))

	stories := make(stories, 0, len(ids))
	storiesCh := make(chan *story, 5)
	for _, id := range ids {
		go func(id int) {
			defer wg.Done()

			url := fmt.Sprintf("%s/v0/item/%v.json", hnUrl, id)
			s, _ := fetchStory(c, url) // TODO: handle error
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
