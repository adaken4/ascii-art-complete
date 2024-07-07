package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"asciiart/ascii"
)

type Ascii struct {
	Result string
	Style  string
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)

	fmt.Println("Starting the Server at port 8080")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request to %s", r.URL.Path)
	if r.Method == http.MethodGet {
		t, err := template.ParseFiles("index.html")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		t.Execute(w, nil)
	} else {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request to %s", r.URL.Path)
	if r.Method == http.MethodPost {
		text := r.FormValue("text")
		banner := r.FormValue("banner")
		result := generateAsciiArt(text, banner)
		t, err := template.ParseFiles("index.html")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		ascii := Ascii{
			Result: result,
		}
		t.Execute(w, ascii)
	} else {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
}

func generateAsciiArt(text, banner string) string {
	// Create a map of each printable Ascii rune to art
	runeAsciiArtMap, err := ascii.RuneAsciiArtMapCreator("./banners/" + banner + ".txt")
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}

	// Define parameters for Ascii art retrieval
	params := ascii.ArtParams{
		InputText:   text,
		SubString:   "",
		Colour:      "",
		AsciiArtMap: runeAsciiArtMap,
	}

	// Build art representation of the input text
	artText, err := ascii.ArtStringBuilder(params)
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
	// Implement the logic to generate ASCII art based on the text and banner
	// This is a placeholder for the actual implementation
	return artText
}
