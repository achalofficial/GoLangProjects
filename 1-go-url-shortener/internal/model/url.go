package model

type ShortenRequest struct {
	URL         string `json:"url" binding:"required,url"`
	CustomAlias string `json:"custom_alias,omitempty"`
	ExpiresAt   string `json:"expires_at,omitempty"` // ISO 8601 optional
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}
