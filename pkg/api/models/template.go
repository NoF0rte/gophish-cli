package models

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v2"
)

// Template models hold the attributes for an e-mail template to be sent to targets
type Template struct {
	ID             int64           `json:"id" yaml:"-"`
	Name           string          `json:"name" yaml:"name"`
	EnvelopeSender string          `json:"envelope_sender" yaml:"-"`
	Subject        string          `json:"subject" yaml:"subject"`
	Text           string          `json:"text,omitempty" yaml:"text,omitempty"`
	TextFile       string          `json:"-" yaml:"text-file,omitempty"`
	HTML           string          `json:"html,omitempty" yaml:"html,omitempty"`
	HTMLFile       string          `json:"-" yaml:"html-file,omitempty"`
	ModifiedDate   time.Time       `json:"modified_date" yaml:"-"`
	Attachments    []Attachment    `json:"attachments" yaml:"-"`
	ProfileFile    string          `json:"-" yaml:"profile-file,omitempty"`
	Profile        *SendingProfile `json:"-" yaml:"profile,omitempty"`
}

func (t *Template) ToJSON() (string, error) {
	data, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (t *Template) replaceVars(vars map[string]string) error {
	name, err := templateReplace(t.Name, vars)
	if err != nil {
		return err
	}
	t.Name = name

	subject, err := templateReplace(t.Subject, vars)
	if err != nil {
		return err
	}
	t.Subject = subject

	if t.Profile != nil {
		err = t.Profile.replaceVars(vars)
	}

	return err
}

func TemplateFromFile(file string, vars map[string]string) (*Template, error) {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var template Template
	err = yaml.Unmarshal(bytes, &template)
	if err != nil {
		return nil, err
	}

	parentDir := filepath.Dir(file)
	if template.Text == "" && template.TextFile != "" {
		textFile := filepath.Join(parentDir, template.TextFile)

		bytes, err = os.ReadFile(textFile)
		if err != nil {
			return nil, err
		}
		template.Text = string(bytes)
	}

	if template.HTML == "" && template.HTMLFile != "" {
		htmlFile := filepath.Join(parentDir, template.HTMLFile)

		bytes, err = os.ReadFile(htmlFile)
		if err != nil {
			return nil, err
		}
		template.HTML = string(bytes)
	}

	if template.Profile == nil && template.ProfileFile != "" {
		profileFile := filepath.Join(parentDir, template.ProfileFile)
		profile, err := SendingProfileFromFile(profileFile, vars)
		if err != nil {
			return nil, err
		}

		template.Profile = profile
	}

	err = template.replaceVars(vars)
	if err != nil {
		return nil, err
	}

	return &template, nil
}

// Attachment contains the fields for an e-mail attachment
type Attachment struct {
	Content string `json:"content"`
	Type    string `json:"type"`
	Name    string `json:"name"`
}
