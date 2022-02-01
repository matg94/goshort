package controller

type URL struct {
	Original  string
	Shortened string
}

type URLConvertRequest struct {
	URL string `json:"url"`
}

type URLCustomConvertRequest struct {
	URL   string `json:"url"`
	Short string `json:"short"`
}
