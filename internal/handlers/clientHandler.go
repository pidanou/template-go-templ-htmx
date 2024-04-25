package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/sevendycom/poc-htmx-alpine/internal/services"
	"github.com/sevendycom/poc-htmx-alpine/internal/views"
)

type ClientHandler struct {
	clientService *services.ClientService
}

func NewClientHandler(clientService *services.ClientService) *ClientHandler {
	return &ClientHandler{
		clientService: clientService,
	}
}

func (h *ClientHandler) GetClients(w http.ResponseWriter, r *http.Request) error {
	clients, err := h.clientService.FindAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return err
	}
	templ.Handler(views.Clients(clients)).ServeHTTP(w, r)
	return nil
}

func (h *ClientHandler) GetClientById(w http.ResponseWriter, r *http.Request) error {
	idString := chi.URLParam(r, "id")
	id, _ := uuid.Parse(idString)
	client, err := h.clientService.FindByID(id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return err
	}

	switch returnType := shouldReturn(r); returnType {
	case "json":
		clientJson, err := json.Marshal(client)
		w.Header().Set("Content-Type", "application/json")
		w.Write(clientJson)
		return err

	case "partial":
		templ.Handler(views.ClientsTableRow(client)).ServeHTTP(w, r)
		return nil

	case "full":
		templ.Handler(views.ClientDetail(client)).ServeHTTP(w, r)
		return nil
	}

	return nil
}

func (h *ClientHandler) GetEditClientById(w http.ResponseWriter, r *http.Request) error {
	idString := chi.URLParam(r, "id")
	id, _ := uuid.Parse(idString)
	client, err := h.clientService.FindByID(id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return err
	}

	switch returnType := shouldReturn(r); returnType {
	case "json":
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
		return nil

	case "partial":
		templ.Handler(views.ClientsTableEditRow(client)).ServeHTTP(w, r)
		return nil

	case "full":
		templ.Handler(views.ClientDetail(client)).ServeHTTP(w, r)
		return nil
	}

	return nil
}
