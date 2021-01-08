package hello

import (
	"encoding/json"
	"net/http"
)

type applicant struct {
	Name     string
	Position string
}

type feedback struct {
	Name     string
	Position string
	Response string
}

func HelloApi(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		response := &feedback{
			Response: "Method not allowed",
		}
		json.NewEncoder(w).Encode(&response)
		return
	}
	token := r.Header.Get("Bearer")

	if len(token) == 0 {
		w.WriteHeader(http.StatusMethodNotAllowed)
		response := &feedback{
			Response: "Bearer token not attched",
		}
		json.NewEncoder(w).Encode(&response)
		return
	}

	var a applicant

	json.NewDecoder(r.Body).Decode(&a)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Token", "my  token")
	w.WriteHeader(http.StatusOK)

	response := &feedback{
		Name:     a.Name,
		Position: a.Position,
	}
	if a.Name == "Marcos Celio" {
		response.Response = "Welcome to the team"
	} else {
		response.Response = "It was not this time. :( Good bye!!"
	}

	// encoding our response to JSON and writing to the responseWriter
	json.NewEncoder(w).Encode(response)

}
