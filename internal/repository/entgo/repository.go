package entgo

import (
	"github.com/DanielTitkov/predictor/internal/repository/entgo/ent"
	"github.com/DanielTitkov/predictor/logger"
)

type EntgoRepository struct {
	client *ent.Client
	logger *logger.Logger
}

func NewEntgoRepository(
	client *ent.Client,
	logger *logger.Logger,
) *EntgoRepository {
	return &EntgoRepository{
		client: client,
		logger: logger,
	}
}
