package controllers

import (
	"strings"
)

type StatusMessage struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func getValidationErrors(verrs map[string][]string) string {
	keys := make([]string, 0, len(verrs))
	for k := range verrs {
		keys = append(keys, k)
	}
	return strings.Join(keys, "\n")
}
