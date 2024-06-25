package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWebApi(t *testing.T) {
	router := createHttpServer()

	t.Run("GET /init", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/init", nil)
		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
		}

		state := toState(t, w)

		if state.MinesLeft != 10 {
			t.Errorf("Expected 10 mines left, got %d", state.MinesLeft)
		}
	})

	t.Run("GET /step/:x/:y/:b", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/init", nil)
		router.ServeHTTP(w, req)
		state := toState(t, w)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
		}

		req, _ = http.NewRequest("GET", "/step/0/0/1", nil)
		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
		}

		state = toState(t, w)

		if state.Steps != 1 {
			t.Errorf("Expected 1 step, got %d", state.Steps)
		}
	})
}

func toState(t *testing.T, w *httptest.ResponseRecorder) GameState {
	var state GameState
	bodyBytes, err := io.ReadAll(w.Body)
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(bodyBytes, &state)
	if err != nil {
		fmt.Println("Error:", err)
		t.Fatal(err)
	}
	return state
}
