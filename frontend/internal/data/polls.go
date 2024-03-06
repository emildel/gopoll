package data

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type Poll struct {
	PollSession string
	Title       string
	Answers     []string
	ExpiresAt   time.Time
}

type PollModelInterface interface {
	Insert(poll *Poll) error
	Get(pollId string) (*Poll, error)
}

type PollModel struct {
	DB *pgxpool.Pool
}

func (p *PollModel) Insert(poll *Poll) error {
	query := `
		INSERT INTO poll (pollSession, title, answers, expires_at)
		VALUES ($1, $2, $3, $4)
		`

	args := []any{poll.PollSession, poll.Title, poll.Answers, poll.ExpiresAt}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()

	if _, err := p.DB.Exec(ctx, query, args...); err != nil {
		return err
	}

	return nil
}

func (p *PollModel) Get(pollId string) (*Poll, error) {
	query := `
		SELECT pollSession, title, answers, expires_at
		FROM poll
		WHERE pollSession = $1
		`

	var poll Poll

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()

	err := p.DB.QueryRow(ctx, query, pollId).Scan(
		&poll.PollSession,
		&poll.Title,
		&poll.Answers,
		&poll.ExpiresAt,
	)

	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &poll, nil
}
