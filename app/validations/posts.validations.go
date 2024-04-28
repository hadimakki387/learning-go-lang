package validations

type CreatePostStruct struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}

type UpdatePostStruct struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
