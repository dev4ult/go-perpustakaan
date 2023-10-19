package helpers

import (
	"github.com/go-playground/validator/v10"
)

func ErrorMapValidation(err error) []map[string]string {
	var mapped = []map[string]string{} 
	for _, err := range err.(validator.ValidationErrors) {
		mapped = append(mapped, map[string]string {
			"field": err.Field(),
			"tag": err.ActualTag(),
		})
	}

	return mapped
}