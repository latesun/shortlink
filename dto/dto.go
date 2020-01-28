package models

type GenerateShortDTO struct {
	LongURL string `json:"url"`
	Secret  string `json:"secret"`
	Length  int    `json:"length"`
	Expires int    `json:"expires"`
}
