package models

type UrlData struct {
	Id       int    `db:"id" json:"id"`
	Url      string `db:"url" json:"url"`
	ShortUrl string `db:"short_url" json:"shortUrl"`
}

type ShortenRequest struct {
	Url string `json:"url"`
}
type Result struct {
	Success  bool   `json:"success"`
	Error    error  `json:"error"`
	ShortUrl string `json:"short_url"`
}
