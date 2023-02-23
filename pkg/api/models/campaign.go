package models

import "encoding/json"

type Campaign struct {
	ID            int64     `json:"id" yaml:"name"`
	Name          string    `json:"name" yaml:"name"`
	CreatedDate   string    `json:"created_date" yaml:"created_date"`
	LaunchDate    string    `json:"launch_date" yaml:"launch_date"`
	SendByDate    string    `json:"send_by_date" yaml:"send_by_date"`
	CompletedDate string    `json:"completed_date" yaml:"completed_date"`
	Template      *Template `json:"template" yaml:"template"`
	//Page          Page     `json:"page" yaml:"page"`
	Status   string          `json:"status" yaml:"status"`
	Results  []*Result       `json:"results" yaml:"results"`
	Groups   []*Group        `json:"groups" yaml:"groups"`
	Timeline []*Event        `json:"timeline" yaml:"timeline"`
	SMTP     *SendingProfile `json:"smtp" yaml:"smtp"`
	URL      string          `json:"url" yaml:"url"`
}

func (c *Campaign) ToJSON() (string, error) {
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "", err
	}

	return string(data), nil
}

type Result struct {
	ID        string  `json:"id" yaml:"id"`
	FirstName string  `json:"first_name" yaml:"first_name"`
	LastName  string  `json:"last_name" yaml:"last_name"`
	Position  string  `json:"position" yaml:"position"`
	Status    string  `json:"status" yaml:"status"`
	IP        string  `json:"ip" yaml:"ip"`
	Latitude  float64 `json:"latitude" yaml:"latitude"`
	Longitude float64 `json:"longitude" yaml:"longitude"`
	SendDate  string  `json:"send_date" yaml:"send_date"`
	Reported  bool    `json:"reported" yaml:"reported"`
}

type Event struct {
	Email   string `json:"email" yaml:"email"`
	Time    string `json:"time" yaml:"time"`
	Message string `json:"message" yaml:"message"`
	Details string `json:"details" yaml:"details"`
}
