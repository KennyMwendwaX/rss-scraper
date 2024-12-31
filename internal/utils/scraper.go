package utils

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/KennyMwendwaX/rss-scraper/internal/config"
	"github.com/KennyMwendwaX/rss-scraper/internal/database"
)

func StartScraping(
	cfg *config.APIConfig,
	concurrency int,
	timeBetweenRequest time.Duration,
) {
	log.Printf("Scraping on %v goroutines every %s duration\n", concurrency, timeBetweenRequest)

	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C {
		feeds, err := cfg.DB.GetNextFeedsToFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Printf("Error getting feeds to fetch: %v", err)
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go scrapeFeed(wg, cfg, feed)
		}

		wg.Wait()
	}
}

func scrapeFeed(wg *sync.WaitGroup, cfg *config.APIConfig, feed database.Feed) {
	defer wg.Done()

	_, err := cfg.DB.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Error marking feed as fetched: %v", err)
		return
	}

	rssFeed, err := UrlToFeed(feed.Url)
	if err != nil {
		log.Printf("Error fetching feed: %v", err)
		return
	}

	for _, item := range rssFeed.Channel.Item {
		log.Println("Found item: ", item.Title, "on feed", feed.Name)
	}

	log.Printf("Feed %s collected, %v post found", feed.Name, len(rssFeed.Channel.Item))
}
