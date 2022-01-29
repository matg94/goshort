package models

type URL struct {
	Original  string
	Shortened string
}

type URLConvertRequest struct {
	URL string `json:"url"`
}
