package repositories

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sevendycom/poc-htmx-alpine/internal/models"
)

type PostgresClientRepo struct {
	Db *sqlx.DB
}

func NewPostgresClientRepo(db *sqlx.DB) *PostgresClientRepo {
	return &PostgresClientRepo{Db: db}
}

func (r *PostgresClientRepo) FindByID(id uuid.UUID) (*models.Client, error) {
	client := models.Client{}
	err := r.Db.Get(&client, "select id, first_name, last_name from public.client where id=$1", id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	return &client, nil
}

func (r *PostgresClientRepo) FindAll() ([]*models.Client, error) {

	var clients []*models.Client

	err := r.Db.Select(&clients, "SELECT * FROM public.client")

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	return clients, nil
}
