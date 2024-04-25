package repositories

import (
	"github.com/google/uuid"
	"github.com/sevendycom/poc-htmx-alpine/internal/models"
)

type ClientRepository interface {
	FindByID(id uuid.UUID) (*models.Client, error)
	FindAll() ([]*models.Client, error)
}
