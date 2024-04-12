package spotify

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

var (
	ClientID        = os.Getenv("Spotify_ID")
	ClientSecret    = os.Getenv("Spotify_Secret")
	spotifyTokenURL = "https://accounts.spotify.com/api/token"
	spotifyAPIURL   = "https://api.spotify.com."
	spotifyAuthURL  = "https://accounts.spotify.com/authorize"
	redirectURL     = "http://localhost:3000"
	Token           AccessToken
	scope           = "user-read-private user-read-email"
	AuthCode        string
	ResponseType    = "authorization_code"
)

type AccessToken struct {
	Token        string `json:"access_token,omitempty"`
	TokenType    string `json:"token_type,omitempty"`
	Duration     uint16 `json:"expires_in,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

func checkSecrets() {
	if ClientID == "" {
		fmt.Println("Client Id is not Set")
		os.Exit(1)
	}

	if ClientSecret == "" {
		fmt.Println("Client Secret is not Set")
		os.Exit(1)
	}
}

func UserAuth() string {
	checkSecrets()

	req, err := http.NewRequest(http.MethodPost, spotifyAuthURL, nil)

	if err != nil {
		fmt.Printf("Error Creating Request for Auth: %v ", err)
		os.Exit(1)
	}

	params := req.URL.Query()
	params.Add("client_id", ClientID)
	params.Add("response_type", "code")
	params.Add("redirect_uri", redirectURL)
	params.Add("scope", scope)

	req.URL.RawQuery = params.Encode()

	fmt.Println(req.URL.String())

	return req.URL.String()

}

func encodeClient(id, secret string) string {
	base := id + ":" + secret
	return base64.StdEncoding.EncodeToString([]byte(base))
}

func GetToken() {
	checkSecrets()

	req, err := http.NewRequest(http.MethodPost, spotifyTokenURL, nil)

	if err != nil {
		fmt.Printf("Error With New Request %v", err)
		os.Exit(1)
	}

	authString := encodeClient(ClientID, ClientSecret)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", authString)

	params := req.URL.Query()
	params.Add("grant_type", "authorization_code")
	params.Add("code", AuthCode)
	params.Add("redirect_uri", redirectURL)
	fmt.Printf("Auth code: %s", AuthCode)

	req.URL.RawQuery = params.Encode()

	client := &http.Client{}

	response, err := client.Do(req)

	if err != nil {
		fmt.Printf("Error Requesting Access Token: %v", err)
		os.Exit(1)
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Printf("Error Reading Body: %v", err)
		os.Exit(1)
	}

	_ = json.Unmarshal(body, &Token)

	fmt.Println(Token)

}

func (a *AccessToken) SetExpirationTime() {}

func (a AccessToken) ExpirationTime() {}

func (a *AccessToken) Refresh() {}
