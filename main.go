package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Snippetbox"))
}

// Add a snippetView handler function.
func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Display a specific snippet... with ID " + strconv.Itoa(id)))
}

// Add a snippetCreate handler function.
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		/*
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("405 - Method Not Allowed"))
		*/
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		log.Print("405 - Method Not Allowed")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Create a new snippet..."))
}
func main() {
	fmt.Println("Finished at 25")
	// Register the two new handler functions and corresponding URL patterns with // the servemux, in exactly the same way that we did before.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)
	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
