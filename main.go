package main

import (
	"log"

	"github.com/vladovidiu/go-podcast/feeds"
	"github.com/vladovidiu/go-podcast/itunes"
)

func main() {
	as := itunes.NewAPIServices()

	res, err := as.Search("Full Stack Radio")

	if err != nil {
		log.Fatalf("error while searching: %v", err)
	}

	for _, item := range res.Results {
		log.Println("---------------------")
		log.Printf("Artist: %s", item.ArtistName)
		log.Printf("Podcast Name: %s", item.TrackName)
		log.Printf("Feed URL: %s", item.FeedURL)

		feed, err := feeds.GetFeed(item.FeedURL)

		if err != nil {
			log.Fatalf("error while getting the feed: %v", err)
		}

		for _, podcast := range feed.Channel.Item {
			log.Println("---------------------")
			log.Printf("Title: %s", podcast.Title)
			log.Printf("Duration: %s", podcast.Duration)
			log.Printf("Description: %s", podcast.Description)
			log.Printf("URL: %s", podcast.Enclosure.URL)
			log.Println("---------------------")
		}

		log.Println("---------------------")
	}
}
