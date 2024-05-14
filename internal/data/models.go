package data

import (
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrRecordNotFound = errors.New("poll not found")
	ErrDuplicateKey   = errors.New("poll session already exists")
)

type Models struct {
	Polls PollModelInterface
}

func NewModel(db *pgxpool.Pool) Models {
	return Models{
		Polls: &PollModel{DB: db},
	}
}
