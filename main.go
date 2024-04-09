/*
Copyright © 2024 Christopher Ritter Christopher.Ritter301@gmail.com
*/
package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/fdaygon/rift/pkg/spotify"
	"github.com/go-chi/chi/v5"
)

func HandleAuth(w http.ResponseWriter, r *http.Request) {
	var authToken string
	var htmlContent string

	authToken = r.URL.Query().Get("code")
	if authToken != "" {

		content := `<html><head><title>Rift</title></head><body>Successfully Logged in</body></html>`
		htmlContent = content
	} else {

		welcomeContent := `<html><head><title>Rift</title></head><body>Welcome to rift. Login in <a href="http://localhost:3000/login">Here<a/> /body></html>`
		htmlContent = welcomeContent
	}
	w.Header().Set("Content-Type", "text/html")

	// Generate the HTML content dynamically

	// Write the HTML content to the response writer
	fmt.Fprint(w, htmlContent)

}

func HandleLogin(w http.ResponseWriter, r *http.Request) {

	url := spotify.UserAuth()

	http.Redirect(w, r, url, http.StatusSeeOther)

}

func main() {
	router := chi.NewRouter()
	//	go commands.Execute()
	router.Get("/", HandleAuth)
	router.Get("/login", HandleLogin)

	exec.Command("open", "http://localhost:3000").Run()

	if err := http.ListenAndServe(":3000", router); err != nil {
		fmt.Println("Unable to start server")
		os.Exit(1)
	}

}
