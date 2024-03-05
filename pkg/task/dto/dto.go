package dto

type Choice string

type CreateTask struct {
	Title  string `json:"title"`
	Status Choice `json:"status"`
}

type UpdateTask struct {
	Title  string `json:"title"`
	Status Choice `json:"status"`
}
