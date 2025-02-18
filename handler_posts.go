package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/mcriq/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	var limit int
	if len(cmd.Args) > 1 {
		return fmt.Errorf("usage: %s [limit]", cmd.Name)
	}

	if  len(cmd.Args) == 1 {
		var err error
		limit, err = strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("invalid lim %s: must be a number", cmd.Args[0])
		}
	} else {
		limit = 2
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit: int32(limit),
	})
	if err != nil {
		return fmt.Errorf("unable to get posts for user '%s': %v", user.Name, err)
	}
	for _, post := range posts {
		fmt.Printf("Title: %s", post.Title)
	}
	return nil
}