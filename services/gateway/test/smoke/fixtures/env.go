package fixtures

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"
)

type RuntimeData struct {
	Password          string
	AdminEmail        string
	AdminPassword     string
	OwnerEmail        string
	ManagedEmail      string
	OwnerUsername     string
	ManagedUsername   string
	HotelTitle        string
	HotelTitleRenamed string

	OwnerAccess  string
	OwnerRefresh string
	AdminAccess  string

	OwnerID   int64
	ManagedID int64
	HotelID   string
	HotelSlug string
	RoomID    string
	BookingID string
}

type Env struct {
	t *testing.T

	baseURL string
	client  *http.Client

	Data RuntimeData
}

func New(t *testing.T) *Env {
	t.Helper()

	ts := time.Now().Unix()
	baseURL := strings.TrimSpace(os.Getenv("GATEWAY_SMOKE_BASE_URL"))
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}
	password := strings.TrimSpace(os.Getenv("GATEWAY_SMOKE_PASSWORD"))
	if password == "" {
		password = "Passw0rd!123"
	}
	adminEmail := strings.TrimSpace(os.Getenv("GATEWAY_SMOKE_ADMIN_EMAIL"))
	if adminEmail == "" {
		adminEmail = "admin@example.com"
	}
	adminPassword := strings.TrimSpace(os.Getenv("GATEWAY_SMOKE_ADMIN_PASSWORD"))
	if adminPassword == "" {
		adminPassword = "Passw0rd!123"
	}

	return &Env{
		t:       t,
		baseURL: strings.TrimRight(baseURL, "/"),
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
		Data: RuntimeData{
			Password:          password,
			AdminEmail:        adminEmail,
			AdminPassword:     adminPassword,
			OwnerEmail:        fmt.Sprintf("http-owner-%d@example.com", ts),
			ManagedEmail:      fmt.Sprintf("http-managed-%d@example.com", ts),
			OwnerUsername:     fmt.Sprintf("http-owner-%d", ts),
			ManagedUsername:   fmt.Sprintf("http-managed-%d", ts),
			HotelTitle:        fmt.Sprintf("HTTP Smoke Hotel %d", ts),
			HotelTitleRenamed: fmt.Sprintf("HTTP Smoke Hotel Renamed %d", ts),
		},
	}
}

func (e *Env) RequestJSON(method, path, bearerToken string, body any, out any) (int, []byte) {
	e.t.Helper()

	var reqBody []byte
	if body != nil {
		var err error
		reqBody, err = json.Marshal(body)
		if err != nil {
			e.t.Fatalf("marshal request body: %v", err)
		}
	}

	req, err := http.NewRequest(method, e.baseURL+path, bytes.NewReader(reqBody))
	if err != nil {
		e.t.Fatalf("create request %s %s: %v", method, path, err)
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if bearerToken != "" {
		req.Header.Set("Authorization", "Bearer "+bearerToken)
	}

	resp, err := e.client.Do(req)
	if err != nil {
		e.t.Fatalf("perform request %s %s: %v", method, path, err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		e.t.Fatalf("read response body %s %s: %v", method, path, err)
	}
	if out != nil && len(respBody) > 0 && resp.StatusCode != http.StatusNoContent {
		if err := json.Unmarshal(respBody, out); err != nil {
			e.t.Fatalf("decode response for %s %s: %v; body=%s", method, path, err, string(respBody))
		}
	}

	return resp.StatusCode, respBody
}

func (e *Env) RequireStatus(got, want int, body []byte) {
	e.t.Helper()
	if got != want {
		e.t.Fatalf("unexpected status: got=%d want=%d body=%s", got, want, string(body))
	}
}
