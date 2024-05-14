package spotify

type TopItems struct {
	Items []Items `json:"items"`
}

type Followers struct {
	Total int `json:"total"`
}

type Artists struct {
	Name string `json:"name"`
}

type Items struct {
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
