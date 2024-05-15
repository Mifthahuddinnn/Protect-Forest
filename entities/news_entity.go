package entities

type NewsSource struct {
	SourceID       string `json:"source_id"`
	SourcePriority int    `json:"source_priority"`
	SourceURL      string `json:"source_url"`
	SourceIcon     string `json:"source_icon"`
}

type NewsArticle struct {
	Title       string      `json:"title"`
	Content     string      `json:"content"`
	Source      *NewsSource `json:"source"`
	PublishedAt string      `json:"pubDate"`
}

type NewsResponse struct {
	Status       string        `json:"status"`
	TotalResults int           `json:"totalResults"`
	Articles     []NewsArticle `json:"results"`
}
