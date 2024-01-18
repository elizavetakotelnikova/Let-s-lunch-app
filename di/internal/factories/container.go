package factories

import (
	"cmd/app/config"
	"cmd/di/internal/lookup"
	"context"
	"database/sql"
	"log"
)

func CreateConfig(ctx context.Context, c lookup.Container) config.Params {
	panic("not implemented")
}

func CreateLogger(ctx context.Context, c lookup.Container) *log.Logger {
	panic("not implemented")
}

func CreateDB(ctx context.Context, c lookup.Container) *sql.DB {
	panic("not implemented")
}
