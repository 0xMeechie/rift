package player

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/fdaygon/rift/pkg/spotify"
)

func GetUser() {
	request, err := http.NewRequest("GET", spotify.SpotifyAPIURL+"/me", nil)
	if err != nil {
		fmt.Println("Error making request")
		os.Exit(1)
	}

	request.Header.Add("Authorization", "Bearer"+spotify.AuthCode)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error Getting user Profile")
	}

	bytebody, _ := io.ReadAll(response.Body)

	fmt.Println(string(bytebody))

}
