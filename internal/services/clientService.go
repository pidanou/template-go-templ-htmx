package services

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/sevendycom/poc-htmx-alpine/internal/models"
	"github.com/sevendycom/poc-htmx-alpine/internal/repositories"
)

type ClientService struct {
	clientRepo repositories.ClientRepository
}

func NewClientService(clientRepo repositories.ClientRepository) *ClientService {
	return &ClientService{clientRepo: clientRepo}
}

func (s *ClientService) FindAll() ([]*models.Client, error) {

	clients, err := s.clientRepo.FindAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return nil, err
	}
	return clients, nil
}

func (s *ClientService) FindByID(id uuid.UUID) (*models.Client, error) {

	client, err := s.clientRepo.FindByID(id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return nil, err
	}
	return client, nil
}
