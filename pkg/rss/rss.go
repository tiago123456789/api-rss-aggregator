package rss

import (
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"time"
)

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	PublishedAt string `xml:"pubDate"`
}

type Feed struct {
	Channel struct {
		Item []Item `xml:"item"`
	} `xml:"channel"`
}

func Parse(url string) Feed {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	response, err := httpClient.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var rssFeed Feed
	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	xml.Unmarshal(data, &rssFeed)

	return rssFeed
}
