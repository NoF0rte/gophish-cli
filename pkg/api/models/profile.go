package models

import (
	"encoding/json"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type InterfaceType string

const (
	InterfaceSMTP InterfaceType = "SMTP"
)

// SendingProfile contains the attributes needed to handle the sending of campaign emails
type SendingProfile struct {
	ID               int64         `json:"id" yaml:"-"`
	Interface        InterfaceType `json:"interface_type" yaml:"-"`
	Name             string        `json:"name" yaml:"name"`
	Host             string        `json:"host" yaml:"host"`
	Username         string        `json:"username,omitempty" yaml:"username,omitempty"`
	Password         string        `json:"password,omitempty" yaml:"password,omitempty"`
	FromAddress      string        `json:"from_address" yaml:"from"`
	IgnoreCertErrors bool          `json:"ignore_cert_errors" yaml:"ignore-cert-errors"`
	Headers          []Header      `json:"headers" yaml:"headers"`
	ModifiedDate     time.Time     `json:"modified_date" yaml:"-"`
	varsReplaced     bool
}

func (s *SendingProfile) ToJSON() (string, error) {
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (s *SendingProfile) replaceVars(vars map[string]string) error {
	if s.varsReplaced {
		return nil
	}

	name, err := templateReplace(s.Name, vars)
	if err != nil {
		return err
	}
	s.Name = name

	host, err := templateReplace(s.Host, vars)
	if err != nil {
		return err
	}
	s.Host = host

	username, err := templateReplace(s.Username, vars)
	if err != nil {
		return err
	}
	s.Username = username

	password, err := templateReplace(s.Password, vars)
	if err != nil {
		return err
	}
	s.Password = password

	from, err := templateReplace(s.FromAddress, vars)
	if err != nil {
		return err
	}
	s.FromAddress = from

	s.varsReplaced = true
	return nil
}

func SendingProfileFromFile(file string, vars map[string]string) (*SendingProfile, error) {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var profile SendingProfile
	err = yaml.Unmarshal(bytes, &profile)
	if err != nil {
		return nil, err
	}

	err = profile.replaceVars(vars)
	if err != nil {
		return nil, err
	}

	return &profile, nil
}

// Header contains the fields and methods for a sending profile to have
// custom headers
type Header struct {
	Key   string `json:"key" yaml:"key"`
	Value string `json:"value" yaml:"value"`
}
