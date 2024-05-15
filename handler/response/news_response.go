package response

type NewsAPIResponse struct {
	Offset    int    `json:"offset"`
	Number    int    `json:"number"`
	Available int    `json:"available"`
	News      []News `json:"news"`
}

type News struct {
	ID            int      `json:"id"`
	Title         string   `json:"title"`
	Text          string   `json:"text"`
	Summary       string   `json:"summary"`
	URL           string   `json:"url"`
	Image         string   `json:"image"`
	PublishDate   string   `json:"publish_date"`
	Authors       []string `json:"authors"`
	Language      string   `json:"language"`
	SourceCountry string   `json:"source_country"`
	Sentiment     float64  `json:"sentiment"`
}
