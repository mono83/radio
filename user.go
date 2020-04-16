package radio

import (
	"errors"
	"strings"
)

// User describes user
type User struct {
	Schema string
	ID     string
}

func (u User) String() string {
	return u.Schema + "://" + u.ID
}

// ParseUser parses user URI into User struct
func ParseUser(s string) (*User, error) {
	chunks := strings.Split(s, "://")
	if len(chunks) != 2 {
		return nil, errors.New("unsupported user format")
	}
	if len(chunks[0]) == 0 {
		return nil, errors.New("schema missing")
	}
	if len(chunks[1]) == 1 {
		return nil, errors.New("user identifier missing")
	}

	return &User{Schema: chunks[0], ID: chunks[1]}, nil
}
