package main

import (
	"context"
	"fmt"
)


func handlerAgg(s *state, cmd command) error {
	testURL := "https://www.wagslane.dev/index.xml"
	feed, err := fetchFeed(context.Background(), testURL)
	if err != nil {
		return fmt.Errorf("unable to fetch from (%s): %v", testURL, err)
	}
	fmt.Printf("%+v\n", feed)
	return nil
}