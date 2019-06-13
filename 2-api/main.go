package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Post Datentyp
type Post struct {
	//Name  //Typ  // Struct Tags (ohne diese wird im JSON der Name genommen)
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// Description ...
func (p *Post) Description() string {
	return fmt.Sprintf("Post %d: %s", p.ID, p.Title)
}

// Anzahl der Beispiel-Beiträge
const numExamplePosts = 3

// Variable mit den Beispiel-Beiträgen
var examplePosts []Post

// Funktion zum Erstellen der Beispiel-Beiträge
func setupExamplePosts() {
	examplePosts = make([]Post, numExamplePosts)
	for i := 0; i < numExamplePosts; i++ {
		examplePosts[i] = Post{i + 1, fmt.Sprintf("Example %d", i+1), "Blablabla"}

		log.Println(examplePosts[i].Description())
	}
}

func main() {
	// Beispiel-Beiträge erstellen
	setupExamplePosts()

	// Handlerfunktion für GET /posts
	http.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Not supported", http.StatusMethodNotAllowed)
		}
		json.NewEncoder(w).Encode(examplePosts)
	})

	// Verbindungen auf Port 8080 annehmen
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
