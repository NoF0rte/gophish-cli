package models

import "encoding/json"

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
