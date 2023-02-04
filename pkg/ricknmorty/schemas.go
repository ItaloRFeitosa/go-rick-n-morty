package ricknmorty

import "time"

type Paginated[T any] struct {
	Info    Info `json:"info"`
	Results []T  `json:"results"`
}

type Info struct {
	Count int     `json:"count"`
	Pages int     `json:"pages"`
	Next  *string `json:"next"`
	Prev  *string `json:"prev"`
}

type Character struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Status   string    `json:"status"`
	Species  string    `json:"species"`
	Type     string    `json:"type"`
	Gender   string    `json:"gender"`
	Origin   Origin    `json:"origin"`
	Location Location  `json:"location"`
	Image    string    `json:"image"`
	Episode  []string  `json:"episode"`
	URL      string    `json:"url"`
	Created  time.Time `json:"created"`
}

type Origin struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PaginatedCharacters Paginated[Character]
