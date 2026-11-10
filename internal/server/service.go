package server

import (
	"context"
	"errors"
	"strings"
	"time"

	db "github.com/IvanSR01/infrapilot/internal/storage/generated"
)

var ErrInvalidInput = errors.New("invalid server data")

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func toResponse(s db.Server) ServerResponse {
	return ServerResponse{
		ID:          s.ID,
		Name:        s.Name,
		Address:     s.Address,
		Environment: s.Environment,
		Status:      s.Status,
		CreatedAt:   s.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt:   s.UpdatedAt.Time.Format(time.RFC3339),
	}
}

func (s *Service) Create(
	ctx context.Context,
	req CreateServerRequest,
) (ServerResponse, error) {
	req.Name = strings.TrimSpace(req.Name)
	req.Address = strings.TrimSpace(req.Address)
	req.Environment = strings.TrimSpace(req.Environment)

	if req.Name == "" || req.Address == "" {
		return ServerResponse{}, ErrInvalidInput
	}

	if len(req.Name) > 100 || len(req.Address) > 255 {
		return ServerResponse{}, ErrInvalidInput
	}

	if req.Environment == "" {
		req.Environment = "development"
	}

	switch req.Environment {
	case "development", "staging", "production":
	default:
		return ServerResponse{}, ErrInvalidInput
	}

	created, err := s.repo.Create(ctx, req)
	if err != nil {
		return ServerResponse{}, err
	}

	return toResponse(created), nil
}

func (s *Service) List(
	ctx context.Context,
) ([]ServerResponse, error) {
	servers, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]ServerResponse, 0, len(servers))

	for _, srv := range servers {
		result = append(result, toResponse(srv))
	}

	return result, nil
}

func (s *Service) GetById(
	ctx context.Context,
	id int64,
) (ServerResponse, error) {
	server, err := s.repo.GetById(ctx, id)
	if err != nil {
		return ServerResponse{}, err
	}

	return toResponse(server), nil
}
