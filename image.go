package sgdb

type Image struct {
	ID        int     `json:"id"`
	Score     int     `json:"score"`
	Style     style   `json:"style"`
	Width     int     `json:"width"`
	Height    int     `json:"height"`
	Nsfw      bool    `json:"nsfw"`
	Humor     bool    `json:"humor"`
	Notes     string  `json:"notes"`
	Mime      mime    `json:"mime"`
	Language  string  `json:"language"`
	URL       string  `json:"url"`
	Thumb     string  `json:"thumb"`
	Lock      bool    `json:"lock"`
	Epilepsy  bool    `json:"epilepsy"`
	Upvotes   int     `json:"upvotes"`
	Downvotes int     `json:"downvotes"`
	Author    *Author `json:"author"`
}
