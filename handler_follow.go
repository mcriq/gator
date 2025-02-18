package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mcriq/gator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <url>", cmd.Name)
	}

	url := cmd.Args[0]

	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("unable to retrieve feed for url (%s): %v", url, err)
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("unable to create feed_follow record: %v", err)
	}

	fmt.Printf("* Feed Name: %s\n", feedFollow.FeedName)
	fmt.Printf("* Username: %s\n", feedFollow.UserName)

	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <url>", cmd.Name)
	}
	url := cmd.Args[0]
	fmt.Printf("URL: %s\n", url)
	fmt.Printf("Name: %s\n", user.Name)
	err := s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		Url: url,
		Name: user.Name,
	})
	if err != nil {
		return fmt.Errorf("unable to delete feed follow: %v", err)
	}
	return nil
}