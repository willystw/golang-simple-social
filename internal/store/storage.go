package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrNotFound          = errors.New("resource not found")
	QueryTimeoutDuration = time.Second * 5
)

type Storage struct {
	Posts interface {
		Create(context.Context, *Post) error
		GetByID(context.Context, int64) (*Post, error)
		DeleteByID(context.Context, int64) error
		Update(context.Context, *Post) error
		GetUserFeeds(context.Context, int64, PaginatedFeedQuery) ([]PostWithMetadata, error)
	}

	Users interface {
		Create(context.Context, *User) error
		GetByID(context.Context, int64) (*User, error)
	}
	Comments interface {
		GetByPostID(ctx context.Context, postID int64) ([]Comment, error)
		Create(context.Context, *Comment) error
	}
	Followers interface {
		Follow(ctx context.Context, userID, followerID int64) error
		UnFollow(ctx context.Context, userID, followerID int64) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts:     &PostStore{db},
		Users:     &UserStore{db},
		Comments:  &CommentStore{db},
		Followers: &FollowerStore{db},
	}
}
