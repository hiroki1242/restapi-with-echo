package model

type Movie struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Release int    `json:"release"`
}
