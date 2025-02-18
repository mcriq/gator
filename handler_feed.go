package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mcriq/gator/internal/database"
)

func handlerListFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeedsAndUsername(context.Background())
	if err != nil {
		return fmt.Errorf("unable to retrieve feeds: %v", err)
	}

	printFeeds(feeds)
	return nil
}

func handlerFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %v <name> <url>", cmd.Name)
	}
	
	name := cmd.Args[0]
	url := cmd.Args[1]

	if name == "" || url == "" {
		return fmt.Errorf("name and URL cannot be empty")
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("unable to generate feed: %v", err)
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("unable to create feed_follow record: %v", err)
	}

	fmt.Println("Feed created successfully:")
	printFeed(feed)
	fmt.Println()
	fmt.Println("=============================================")

	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* UserID:        %s\n", feed.UserID)
}

func printFeeds(feeds []database.GetFeedsAndUsernameRow) {
	for _, feed := range feeds {
		fmt.Printf("%s\n", feed.Name)
		fmt.Printf("%s\n", feed.Url)
		fmt.Printf("Created by: %s\n", feed.UserName)
		fmt.Println()
	}
}
