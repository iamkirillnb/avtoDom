package entities

type Url struct {
	IncommingUrl string `json:"incomming_url" db:"incomming_url"`
	OutUrl       string `json:"out_url" db:"out_url"`
	Code         int    `json:"code" db:"code"`
}
