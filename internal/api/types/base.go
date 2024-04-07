package types

type PagedResourceList struct {
	Count    int             `json:"count"`
	Next     *string         `json:"next"`
	Previous *string         `json:"previous"`
	Results  []NamedResource `json:"results"`
}

type NamedResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type APIResource struct {
	URL string `json:"url"`
}
