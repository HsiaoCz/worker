package models

type User struct {
	Username string `json:"username"`
	Today    string `json:"today"`
	Luck     Lucky  `json:"luck"`
}

type Lucky struct {
	Value   int    `json:"value"`
	Level   string `json:"level"`
	Explain string `json:"explain"`
}

type Post struct {
	Username    string `json:"username"`
	RequestText string `json:"request_text"`
}
