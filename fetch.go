package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func fetchItemsIds(c *http.Client, url string) ([]int, error) {
	resp, err := c.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

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
