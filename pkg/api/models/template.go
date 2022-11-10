package models

import (
	"encoding/json"
	"time"
)

// Template models hold the attributes for an e-mail template to be sent to targets
type Template struct {
	Id             int64        `json:"id" yaml:"-"`
	Name           string       `json:"name" yaml:"name"`
	EnvelopeSender string       `json:"envelope_sender" yaml:"-"`
	Subject        string       `json:"subject" yaml:"subject"`
	Text           string       `json:"text,omitempty" yaml:"text,omitempty"`
	TextFile       string       `json:"-" yaml:"text-file,omitempty"`
	HTML           string       `json:"html,omitempty" yaml:"html,omitempty"`
	HTMLFile       string       `json:"-" yaml:"html-file,omitempty"`
	ModifiedDate   time.Time    `json:"modified_date" yaml:"-"`
	Attachments    []Attachment `json:"attachments" yaml:"-"`
}

func (t *Template) ToJson() (string, error) {
	data, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// Attachment contains the fields for an e-mail attachment
type Attachment struct {
	Content string `json:"content"`
	Type    string `json:"type"`
	Name    string `json:"name"`
}
