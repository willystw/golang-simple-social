package store

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrDuplicateEmail    = errors.New("a user with that email already exists")
	ErrDuplicateUsername = errors.New("a user with that username already exists")
)

type User struct {
	ID        int64    `json:"id"`
	Username  string   `json:"username"`
	Email     string   `json:"email"`
	Password  password `json:"-"`
	CreatedAt string   `json:"created_at"`
}

type password struct {
	text *string
	hash []byte
}

func (p *password) Set(text string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	p.text = &text
	p.hash = hash
	return nil
}

func (u *User) SetPassword(plaintext string) error {
	if plaintext == "" {
		return errors.New("password cannot be empty")
	}
	u.Password.Set(plaintext)
	return nil
}

type UserStore struct {
	db *sql.DB
}

func (s *UserStore) Create(ctx context.Context, tx *sql.Tx, user *User) error {
	query := `
	INSERT INTO users(username, password, email)
	VALUES ($1, $2, $3) RETURNING id,created_at	
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := tx.QueryRowContext(
		ctx,
		query,
		user.Username,
		user.Password,
		user.Email,
	).Scan(
		&user.ID,
		&user.CreatedAt,
	)

	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return ErrDuplicateEmail
		case err.Error() == `pq: duplicate key value violates unique constraint "users_username_key"`:
			return ErrDuplicateUsername
		default:
			return err
		}
	}

	return nil
}

func (s *UserStore) GetByID(ctx context.Context, userID int64) (*User, error) {
	query := `
		SELECT id, username, email, password, created_at FROM users
		WHERE id=$1

	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	var user User
	err := s.db.QueryRowContext(ctx, query, userID).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}

func (s *UserStore) CreateAndInvite(ctx context.Context, user *User, token string, invitationExp time.Duration) error {
	//tx wrapper
	return withTx(s.db, ctx, func(tx *sql.Tx) error {
		//create user

		if err := s.Create(ctx, tx, user); err != nil {
			return err
		}

		//create user invites
		if err := s.createUserInvitation(ctx, tx, token, invitationExp, user.ID); err != nil {
			return err
		}

		return nil
	})

}

func (s *UserStore) createUserInvitation(ctx context.Context, tx *sql.Tx, token string, invitationExp time.Duration, id int64) error {
	query := `
	Insert into user_invitations(token,user_id,expiry) values ($1, $2, $3)
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := tx.ExecContext(ctx, query, token, id, time.Now().Add(invitationExp))
	if err != nil {
		return err
	}
	return nil
}
