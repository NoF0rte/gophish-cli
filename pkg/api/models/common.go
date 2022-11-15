package models

import (
	"bytes"
	"encoding/json"
	"text/template"
)

type GenericResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
	Data    string `json:"data"`
}

func (r *GenericResponse) ToJson() (string, error) {
	data, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func templateReplace(text string, vars map[string]string) (string, error) {
	t, err := template.New("replacement").Option("missingkey=error").Parse(text)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	err = t.Execute(buf, vars)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
