package models

import "time"

// ArticleList structure
type ArticleList struct {
	Status   string `json:"status"`
	Source   string `json:"source"`
	SortBy   string `json:"sortBy"`
	Articles []struct {
		Author      string    `json:"author"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		URL         string    `json:"url"`
		URLToImage  string    `json:"urlToImage"`
		PublishedAt time.Time `json:"publishedAt"`
	} `json:"articles"`
}

// MediaList structure
type MediaList struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Category    string `json:"category"`
	Language    string `json:"language"`
	Country     string `json:"country"`
	UrlsToLogos struct {
		Small  string `json:"small"`
		Medium string `json:"medium"`
		Large  string `json:"large"`
	} `json:"urlsToLogos"`
	SortBysAvailable []string `json:"sortBysAvailable"`
}
