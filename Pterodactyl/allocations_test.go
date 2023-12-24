package Pterodactyl

import (
	_ "bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockHandler struct{}

func (m *mockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/nodes/1/allocations" {
		_, err := fmt.Fprintln(w, `{"allocations": [{"id":1, "ip":"192.168.1.1"}]}`)
		if err != nil {
			return
		}
	} else {
		http.Error(w, "Error occurred", http.StatusInternalServerError)
	}
}

func TestGetAllocations(t *testing.T) {
	server := httptest.NewServer(&mockHandler{})
	defer server.Close()

	config := &PterodactylConfig{
		URL:    server.URL,
		ApiKey: "test-api-key",
	}

	tests := []struct {
		nodeId  int
		wantLen int
		wantErr error
	}{
		{1, 1, nil},
		{2, 0, errors.New("Error occurred")},
	}

	for _, tt := range tests {
		got, err := config.GetAllocations(tt.nodeId)
		if len(got) != tt.wantLen {
			t.Errorf("Expected allocations length %v, got %v", tt.wantLen, len(got))
		}
		if (err != nil) != (tt.wantErr != nil) {
			t.Errorf("Expected error %v, got %v", tt.wantErr, err)
		}
	}
}
