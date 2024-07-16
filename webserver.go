package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"

	"asciiart/ascii"
)

type Ascii struct {
	Result string
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/ascii-art", asciiArtHandler)

	fmt.Println("Starting the Server at port 8080")
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request to %s\n", r.URL.Path)
	if r.URL.Path != "/" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
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
	fmt.Printf("Received request to %s\n", r.URL.Path)
	if r.URL.Path != "/ascii-art" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
	if r.Method == http.MethodPost {
		text := r.FormValue("text")
		if strings.Contains(text, "\r\n") {
			text = strings.ReplaceAll(text, "\r\n", "\n")
		}
		banner := r.FormValue("banner")
		result := generateAsciiArt(text, banner)
		t, err := template.ParseFiles("result.html")
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
