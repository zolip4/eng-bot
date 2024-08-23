package services

import "engbot/internal/models"

func GetExampleData() models.ExampleModel {
	return models.ExampleModel{
		ID:    1,
		Title: "This is an example",
	}
}
