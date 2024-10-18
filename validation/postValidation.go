package validation

type ValidatePostInput struct {
	Title   string `json:"title" validate:"required,min=3,max=100"`
	Content string `json:"content" validate:"required,min=3"`
}

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
