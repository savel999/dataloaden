//go:generate go run github.com/savel999/dataloaden UserLoader string *github.com/savel999/dataloaden/example.User

package example

import (
	"context"
	"time"
)

// User is some kind of database backed model
type User struct {
	ID   string
	Name string
}

// NewLoader will collect user requests for 2 milliseconds and send them as a single batch to the fetch func
// normally fetch would be a database call.
func NewLoader() *UserLoader {
	return &UserLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(fCtx context.Context, keys []string) ([]*User, []error) {
			users := make([]*User, len(keys))
			errors := make([]error, len(keys))

			for i, key := range keys {
				users[i] = &User{ID: key, Name: "user " + key}
			}
			return users, errors
		},
	}
}
