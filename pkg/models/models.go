package models

type Name struct {
    Name    string  `json:"name"`
    Amount  int `json:"amount"`
}

type Names struct {
    Names []Name `json:"names"`
}
