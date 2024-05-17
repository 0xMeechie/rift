package web

import (
	"fmt"
	"net/http"
	"os"

	"github.com/fdaygon/rift/pkg/spotify"
	"github.com/go-chi/chi/v5"
)

// Need a server to allow us to request an auth code from spotify
func StartServer() {
	router := chi.NewRouter()
	router.Get("/", HandleAuth)
	router.Get("/login", HandleLogin)
	router.Get("/callback", HandleCallBack)

	if err := http.ListenAndServe(":3000", router); err != nil {
		fmt.Println("Unable to start server")
		os.Exit(1)
	}

}

func HandleAuth(w http.ResponseWriter, r *http.Request) {
	var htmlContent string

	spotify.AuthCode = r.URL.Query().Get("code")
	if spotify.AuthCode != "" {
		spotify.GetToken()
		http.Redirect(w, r, "/login", http.StatusSeeOther)

	} else {

		welcomeContent := `<html><head><title>Rift</title></head><body>Welcome to rift. Login in <a href="http://localhost:3000/login">Here<a/> /body></html>`
		htmlContent = welcomeContent
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, htmlContent)
	}

}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	successContent := `<html><head><title>Rift</title></head><body>Successfully Authenticated. You may close out this window!<a/> </body></html>`
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, successContent)

}

func HandleCallBack(w http.ResponseWriter, r *http.Request) {
	//not sure how i want to handle the call back so for now the app will just close out.
	fmt.Println("Log in failed. Closing application")
	os.Exit(1)
}
