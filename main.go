package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Gist struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

type Response struct {
	User  string `json:"user"`
	Gists []Gist `json:"gists"`
}

func getUserGists(user string) ([]Gist, error) {

	url := fmt.Sprintf("https://api.github.com/users/%s/gists", user)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("user not found")
	}

	var data []map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	var gists []Gist

	for _, g := range data {

		id, _ := g["id"].(string)
		url, _ := g["html_url"].(string)

		desc := ""
		if g["description"] != nil {
			if d, ok := g["description"].(string); ok {
				desc = d
			}
		}

		gists = append(gists, Gist{
			ID:          id,
			Description: desc,
			URL:         url,
		})
	}

	return gists, nil
}

func gistHandler(w http.ResponseWriter, r *http.Request) {

	user := r.URL.Path[1:]

	if user == "" {
		http.Error(w, "User required", http.StatusBadRequest)
		return
	}

	gists, err := getUserGists(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	resp := Response{
		User:  user,
		Gists: gists,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})
}

func main() {

	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/", gistHandler)

	fmt.Println("Server running on port 8080")

	http.ListenAndServe(":8080", nil)
}
