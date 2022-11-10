package models

import (
	"encoding/json"
	"time"
)

// Template models hold the attributes for an email template to be sent to targets
type Template struct {
	Id             int64        `json:"id"`
	Name           string       `json:"name" yaml:"name"`
	EnvelopeSender string       `json:"envelope_sender"`
	Subject        string       `json:"subject" yaml:"subject"`
	Text           string       `json:"text,omitempty" yaml:"text"`
	TextFile       string       `json:"-" yaml:"text-file"`
	HTML           string       `json:"html,omitempty" yaml:"html"`
	HTMLFile       string       `json:"-" yaml:"html-file"`
	ModifiedDate   time.Time    `json:"modified_date"`
	Attachments    []Attachment `json:"attachments"`
}

func (t *Template) ToJson() (string, error) {
	data, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// Attachment contains the fields for an email attachment
type Attachment struct {
	Content string `json:"content"`
	Type    string `json:"type"`
	Name    string `json:"name"`
}
