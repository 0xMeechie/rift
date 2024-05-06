package spotify

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type TopParams struct {
	Type      string
	TimeRange string
	Limit     int8
}

func GetAccount() {
	client := http.Client{}

	request, err := http.NewRequest(http.MethodGet, SpotifyAPIURL+"/me", nil)
	if err != nil {
		fmt.Println("error with user request")
	}
	request.Header.Add("Authorization", "Bearer "+RequestToken)
	reponse, err := client.Do(request)
	if err != nil {
		fmt.Println("Error getting user profile")
	}
	byteBody, err := io.ReadAll(reponse.Body)

	if err != nil {
		fmt.Printf("Error Reading Body: %v", err)
		os.Exit(1)
	}

	fmt.Println(string(byteBody))
}

func GetTopItems() {

	client := http.Client{}

	request, err := http.NewRequest(http.MethodGet, SpotifyAPIURL+"me/top/artists", nil)
	if err != nil {
		fmt.Println("Error with request for artists")
	}

}
