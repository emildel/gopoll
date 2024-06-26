package data

import (
	"context"
	"errors"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type Poll struct {
	PollSession string
	Title       string
	Answers     []string
	Results     []int
	ExpiresAt   time.Time
}

type PollModelInterface interface {
	Insert(poll *Poll) error
	Get(pollId string) (*Poll, error)
	Update(pollAnswer int, pollId string) (*Poll, error)
}

type PollModel struct {
	DB *pgxpool.Pool
}

// Insert a new Poll to the database
func (p *PollModel) Insert(poll *Poll) error {
	query := `
		INSERT INTO poll (pollSession, title, answers, results, expires_at)
		VALUES ($1, $2, $3, $4, $5)
		`

	args := []any{poll.PollSession, poll.Title, poll.Answers, poll.Results, poll.ExpiresAt}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var psqlError *pgconn.PgError

	if _, err := p.DB.Exec(ctx, query, args...); err != nil {
		switch {
		case errors.As(err, &psqlError) && psqlError.Code == pgerrcode.UniqueViolation:
			return ErrDuplicateKey
		default:
			return err
		}
	}

	return nil
}

// Get a Poll by poll id.
func (p *PollModel) Get(pollId string) (*Poll, error) {
	query := `
		SELECT pollSession, title, answers, results, expires_at
		FROM poll
		WHERE pollSession = $1
		`

	var poll Poll

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := p.DB.QueryRow(ctx, query, pollId).Scan(
		&poll.PollSession,
		&poll.Title,
		&poll.Answers,
		&poll.Results,
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

func (p *PollModel) Update(pollAnswer int, pollId string) (*Poll, error) {

	query := "CALL poll_updates_results($1, $2, $3, $4)"

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()

	var poll Poll

	err := p.DB.QueryRow(ctx, query, pollId, pollAnswer, nil, nil).Scan(
		&poll.Answers,
		&poll.Results,
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
