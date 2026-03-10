package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"entgo.io/ent/dialect/sql"
	"github.com/echo-logger/service/ent"
	"github.com/echo-logger/service/ent/requestlog"
)

// EchoHandler handles incoming HTTP requests, logs them to the database, and echoes back request info.
type EchoHandler struct {
	client *ent.Client
}

// NewEchoHandler creates a new EchoHandler with the given Ent client.
func NewEchoHandler(client *ent.Client) *EchoHandler {
	return &EchoHandler{client: client}
}

// EchoResponse is the JSON structure returned to the caller.
type EchoResponse struct {
	Method  string              `json:"method"`
	Path    string              `json:"path"`
	Headers map[string][]string `json:"headers"`
	Body    string              `json:"body"`
	IP      string              `json:"ip"`
}

// ServeHTTP handles all incoming requests.
func (h *EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	headersJSON, _ := json.Marshal(r.Header)
	ip := r.RemoteAddr

	// Save request log to database.
	_, err = h.client.RequestLog.Create().
		SetMethod(r.Method).
		SetPath(r.URL.Path).
		SetHeaders(string(headersJSON)).
		SetBody(string(body)).
		SetIP(ip).
		Save(r.Context())
	if err != nil {
		log.Printf("failed to save request log: %v", err)
	}

	resp := EchoResponse{
		Method:  r.Method,
		Path:    r.URL.Path,
		Headers: r.Header,
		Body:    string(body),
		IP:      ip,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("failed to encode response: %v", err)
	}

	fmt.Printf("[%s] %s %s from %s\n", r.Method, r.URL.Path, r.Proto, ip)
}

// ListRequests returns all logged requests from the database ordered by newest first.
func (h *EchoHandler) ListRequests(w http.ResponseWriter, r *http.Request) {
	logs, err := h.client.RequestLog.Query().
		Order(requestlog.ByCreatedAt(sql.OrderDesc())).
		All(r.Context())
	if err != nil {
		http.Error(w, "failed to query request logs", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(logs)
}
