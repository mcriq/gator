package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/mcriq/gator/internal/database"
)


func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <time_between_reqs>", cmd.Name)
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("unable to parse to duration: %v", err)
	}

	fmt.Printf("Collecting feeds every %v\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		if err := scrapeFeeds(s); err != nil {
			fmt.Printf("error scraping feeds: %v", err)
		}
	}
}

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("unable to get next feed: %v", err)
	}

	err = s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		LastFetchedAt: sql.NullTime{
			Time: time.Now().UTC(),
			Valid: true,
		},
		ID: nextFeed.ID,
	})
	if err != nil {
		return fmt.Errorf("unable to mark %s as fetched: %v", nextFeed.Name, err)
	}

	feed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return fmt.Errorf("unable to fetch feed at %s: %v", nextFeed.Url, err)
	}

	fmt.Printf("Fetched Feed: %s\n", nextFeed.Name)
	for _, item := range feed.Channel.Item {
		dateLayout := "Mon, 02 Jan 2006 15:04:05 MST"
		parsedTime, err := time.Parse(dateLayout, item.PubDate)
		if err != nil {
			return fmt.Errorf("unable to parse publish date: %v", err)
		}
		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID: uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Title: item.Title,
			Url: item.Link,
			Description: sql.NullString{
				String: item.Description,
				Valid: true,
			},
			PublishedAt: parsedTime,
			FeedID: nextFeed.ID,
		})
		if err != nil {
			if isUniqueConstraintError(err) {
				continue
			}
			return fmt.Errorf("unable to add feed %s: %v", item.Title, err)
		}
		fmt.Printf("* Title: %s\n", item.Title)
	}
	return nil
}

func isUniqueConstraintError(err error) bool {
	var pqErr *pq.Error
	if errors.As(err, &pqErr) {
		return pqErr.Code == "23505"
	}
	return false
}