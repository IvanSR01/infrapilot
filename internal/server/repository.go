package server

import (
	"context"

	db "github.com/IvanSR01/infrapilot/internal/storage/generated"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	queries *db.Queries
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{
		queries: db.New(pool),
	}
}

func (r *Repository) Create(ctx context.Context, req CreateServerRequest) (db.Server, error) {
	return r.queries.CreateServer(ctx, db.CreateServerParams{
		Name:        req.Name,
		Address:     req.Address,
		Environment: req.Environment,
	})
}

func (r *Repository) List(ctx context.Context) ([]db.Server, error) {
	return r.queries.GetListServers(ctx)
}

func (r *Repository) GetById(ctx context.Context, id int64) (db.Server, error) {
	return r.queries.GetServer(ctx, id)
}
