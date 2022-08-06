package hn

type Story struct {
	ID          int    `json:"id"`                    // unique item's id
	By          string `json:"by,omitempty"`          // username of the item's author
	Descendants int    `json:"descendants,omitempty"` // total comments count
	Kids        []int  `json:"kids,omitempty"`        // ids of item's comments
	Score       int    `json:"score,omitempty"`       // story's score
	Time        int    `json:"time"`                  // creation date of the item
	Title       string `json:"title"`                 // story's title
	Type        string `json:"type"`                  // item's type
	URL         string `json:"url"`                   // url of the story
}

type Stories []*Story

func (s Stories) Len() int { return len(s) }

func (s Stories) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s Stories) Less(i, j int) bool {
	return s[i].ID < s[j].ID
}
