package validation

type ValidatePostInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type ErrorMsg struct {
	Field	string	`json:"field"`
	Message	string	`json:"message"`
}