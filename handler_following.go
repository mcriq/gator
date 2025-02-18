package main

import (
	"context"
	"fmt"

	"github.com/mcriq/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {

	follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("unable to get follows for user: %s", err)
	}

	for _, follow := range follows {
		fmt.Printf("* Feed Name: %s\n", follow.FeedName)
	}
	return nil
}

