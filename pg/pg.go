package pg

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

var (
	pgHost = os.Getenv("PG_HOST")
	pgPort = os.Getenv("PG_PORT")
	pgUser = os.Getenv("PG_USER")
	pgPass = os.Getenv("PG_PASS")
	pgName = os.Getenv("PG_NAME")
)

type Client struct {
	pool *pgxpool.Pool
}

func NewClient(ctx context.Context) (cli *Client, err error) {
	pool, err := pgxpool.New(ctx, fmt.Sprintf("postgres://%s:%s@%s:%s/%s", pgUser, pgPass, pgHost, pgPort, pgName))
	if err != nil {
		log.Warn().Err(err).Msg("pgxpool.New error")
		return nil, err
	}
	return &Client{
		pool: pool,
	}, nil
}
