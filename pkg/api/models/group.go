package models

type Group struct {
	ID           int64    `json:"id" yaml:"id"`
	Name         string   `json:"name" yaml:"name"`
	Targets      []Target `json:"targets" yaml:"targets"`
	ModifiedDate string   `json:"modified_date" yaml:"modified_date"`
}

type Target struct {
	Email     string `json:"email" yaml:"email"`
	FirstName string `json:"first_name" yaml:"first_name"`
	LastName  string `json:"last_name" yaml:"last_name"`
	Position  string `json:"position" yaml:"position"`
}
