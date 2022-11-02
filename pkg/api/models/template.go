package models

import "time"

// Template models hold the attributes for an email template to be sent to targets
type Template struct {
	Id             int64        `json:"id"`
	Name           string       `json:"name"`
	EnvelopeSender string       `json:"envelope_sender"`
	Subject        string       `json:"subject"`
	Text           string       `json:"text"`
	HTML           string       `json:"html"`
	ModifiedDate   time.Time    `json:"modified_date"`
	Attachments    []Attachment `json:"attachments"`
}

// Attachment contains the fields for an email attachment
type Attachment struct {
	Content string `json:"content"`
	Type    string `json:"type"`
	Name    string `json:"name"`
}
