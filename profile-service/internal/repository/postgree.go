package repository

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)

const (
	defaultTimeoutQuery = 5 * time.Second
)

type Repository struct {
	DB     *sql.DB
	Ctx    context.Context
	Cancel context.CancelFunc
}

func NewRepository(ctx context.Context) (*Repository, error) {
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URI"))
	if err != nil {
		return nil, fmt.Errorf("could not connect to the database: %v", err)
	}

	repoCtx, repoCancel := context.WithCancel(ctx)
	return &Repository{
		DB:     db,
		Ctx:    repoCtx,
		Cancel: repoCancel,
	}, nil
}

func (r *Repository) Close() {
	r.DB.Close()
	r.Cancel()
}
