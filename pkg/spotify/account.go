package spotify

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type TopParams struct {
	Type      string
	TimeRange string
	Limit     string
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

func GetTopItems(params TopParams) {
	var Tops TopItems
	if params.Type == "" {
		fmt.Println("Please Select Artist or Tracks")
		os.Exit(1)
	}

	//check to see if there is a limit that is smaller or equal to 0. if it is set to 0 then we are
	//default to 5
	//if params.Limit < 0 {
	//	fmt.Println("Limit must be higher than 0 (1-50)")
	//	os.Exit(1)
	//} else if params.Limit == 0 {
	//	params.Limit = 5
	//}

	// if this isn't set then it will default to medium_term
	if params.TimeRange == "" {
		params.TimeRange = "medium_term"
	}

	client := http.Client{}

	request, err := http.NewRequest(http.MethodGet, SpotifyAPIURL+"/me/top/"+params.Type, nil)
	if err != nil {
		fmt.Println("Error with request for artists")
	}

	reqParams := request.URL.Query()
	reqParams.Add("time_range", params.TimeRange)
	reqParams.Add("limit", params.Limit)

	request.URL.RawQuery = reqParams.Encode()

	request.Header.Add("Authorization", "Bearer "+RequestToken)

	response, err := client.Do(request)

	if err != nil {
		fmt.Println("Error request tops: ", err)
		os.Exit(1)

	}

	if response.StatusCode == 401 {
		fmt.Println("Token is Expired. Please Login and refresh")
		return
	}

	btyebody, _ := io.ReadAll(response.Body)

	json.Unmarshal(btyebody, &Tops)
	DisplayTops(Tops)

}

func DisplayTops(topItems TopItems) {
	fmt.Printf("Here are your top %s\n", topItems.Items[0].Type)

	if topItems.Items[0].Type == "track" {
		for index, top := range topItems.Items {
			fmt.Printf("%v. %s by %s\n", index+1, top.Name, top.Artists[0].Name)
		}
	} else {
		for index, top := range topItems.Items {
			fmt.Printf("%v. %s\n", index+1, top.Name)
		}
	}

}

func GetPlaylist() {

}
