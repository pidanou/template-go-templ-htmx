package api

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/sevendycom/poc-htmx-alpine/internal/datastore"
	"github.com/sevendycom/poc-htmx-alpine/internal/handlers"
	"github.com/sevendycom/poc-htmx-alpine/internal/repositories"
	"github.com/sevendycom/poc-htmx-alpine/internal/services"
)

type Server struct {
	port string
}

func NewServer(port string) *Server {
	return &Server{port: port}
}

func (s *Server) Start() error {
	postgresDb := datastore.NewPostgresDb(os.Getenv("POSTGRES_URL"))

	r := chi.NewRouter()

	fileServer(r, "/static", "static")

	clientRepo := repositories.NewPostgresClientRepo(postgresDb)
	clientService := services.NewClientService(clientRepo)
	clientHandler := handlers.NewClientHandler(clientService)
	r.Get("/clients", handle(clientHandler.GetClients))
	r.Get("/clients/{id}", handle(clientHandler.GetClientById))
	r.Get("/clients/edit/{id}", handle(clientHandler.GetEditClientById))

	return http.ListenAndServe(s.port, r)
}
