package model

type User struct {
	Name     string `json:"Name"`
	LastName string `json:"LastName"`
	SurName  string `json:"SurName"`
	Gender   string `json:"Gender"`
	Status   string `json:"Status"`
}
