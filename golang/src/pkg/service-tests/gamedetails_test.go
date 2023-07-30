package it_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	games "gitlab.com/test-group3671509/source/go-lang-service/pkg/service-tests/httpserver"
)

func TestGETGameDetails(t *testing.T) {
	t.Run("returns Pepper's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()
		games.GameServer(response, request)
		want := "Game"
		got := response.Body.String()
		if strings.Contains(got, want) != true {
			t.Errorf("Expected want %q, Received got %q", want, got)
		}
	})
}
