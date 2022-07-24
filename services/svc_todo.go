package services

import "web-api/model"

func TodoTemplate() *model.TodoPageData {
	data := model.TodoPageData{
		PageTitle: "My TODO list",
		Todos: []model.Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}
	return &data
}
