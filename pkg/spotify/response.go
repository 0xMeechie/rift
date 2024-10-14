package spotify

type Top struct {
	Items []TopItems `json:"items"`
}

type Followers struct {
	Total int `json:"total"`
}

type Artists struct {
	Name string `json:"name"`
}

type TopItems struct {
	Followers Followers `json:"followers"`
	Genres    []string  `json:"genres"`
	Href      string    `json:"href"`
	ID        string    `json:"id"`
	Artists   []Artists `json:"artists"`

	Name       string `json:"name"`
	Popularity int    `json:"popularity"`
	Type       string `json:"type"`
	URI        string `json:"uri"`
}

type Playlist struct {
	Href  string          `json:"href"`
	Next  string          `json:"next"`
	Total int             `json:"total"`
	Items []PlaylistItems `json:"items"`
}
type ExternalUrls struct {
	Spotify string `json:"spotify"`
}
type Owner struct {
	ExternalUrls ExternalUrls `json:"external_urls"`
	Href         string       `json:"href"`
	ID           string       `json:"id"`
	Type         string       `json:"type"`
	URI          string       `json:"uri"`
	DisplayName  string       `json:"display_name"`
}
type Tracks struct {
	Href  string `json:"href"`
	Total int    `json:"total"`
}
type PlaylistItems struct {
	Collaborative bool   `json:"collaborative"`
	Description   string `json:"description"`
	ID            string `json:"id"`
	Name          string `json:"name"`
	Owner         Owner  `json:"owner"`
	Public        bool   `json:"public"`
	Tracks        Tracks `json:"tracks"`
	Type          string `json:"type"`
}
