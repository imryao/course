package pg

import (
	"context"
	"github.com/imryao/course/types"
	"github.com/rs/zerolog/log"
	"time"
)

func (c *Client) InsertUser(ctx context.Context, user *types.User) (err error) {
	now := time.Now()
	_, err = c.pool.Exec(ctx, `INSERT INTO "public"."course_users" ("status", "uid", "name", "email", "inserted_at", "updated_at") VALUES (0, $1, $2, $3, $4, $5) ON CONFLICT("uid") DO UPDATE SET "name"=EXCLUDED.name, email=EXCLUDED.email, updated_at=EXCLUDED.updated_at;`, *user.Uid, *user.Name, *user.Email, now, now)
	if err != nil {
		log.Warn().Int64("uid", *user.Uid).Err(err).Msg("pg.InsertUser c.pool.Exec error")
		return err
	}
	log.Info().Int64("uid", *user.Uid).Msg("pg.InsertUser ok")
	return nil
}
