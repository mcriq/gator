package main

import (
	"context"
	"fmt"
)


func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("unable to reset users table: %v", err)
	}
	fmt.Printf("%v command successful: users table has been emptied\n", cmd.Name)
	return nil
}