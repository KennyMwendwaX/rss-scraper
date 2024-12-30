package utils

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Lang        string    `xml:"language"`
		Items       []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func urlToFeed(url string) (RSSFeed, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	response, err := httpClient.Get(url)
	if err != nil {
		return RSSFeed{}, err
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return RSSFeed{}, err
	}
	feed := RSSFeed{}
	err = xml.Unmarshal(data, &feed)
	if err != nil {
		return RSSFeed{}, err
	}
	return feed, nil
}
