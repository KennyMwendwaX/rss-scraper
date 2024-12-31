package utils

import (
	"context"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/KennyMwendwaX/rss-scraper/internal/config"
	"github.com/KennyMwendwaX/rss-scraper/internal/database"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
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
		pubAt, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Printf("Couldn't parse date %v with err %v", item.PubDate, err)
			continue
		}

		id := uuid.New()
		pgID := pgtype.UUID{
			Bytes: id,
			Valid: true,
		}

		now := time.Now().UTC()
		pgTimestamp := pgtype.Timestamp{
			Time:  now,
			Valid: true,
		}

		publishedAt := pgtype.Timestamp{
			Time:  pubAt,
			Valid: true,
		}

		_, err = cfg.DB.CreatePost(context.Background(), database.CreatePostParams{
			ID:          pgID,
			Title:       item.Title,
			Description: item.Description,
			Url:         item.Link,
			PublishedAt: publishedAt,
			CreatedAt:   pgTimestamp,
			UpdatedAt:   pgTimestamp,
			FeedID:      feed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			log.Printf("Failed to create post: %v", err)
		}
	}

	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))
}
