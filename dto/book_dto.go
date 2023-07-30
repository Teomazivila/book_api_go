package dto

type BookCreatedDTO struct {
	Title       string `json:"title" form:"title" binding:"required"`
	Year        string `json:"year" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	BookCover   string `json:"book_cover" form:"book_cover" `
	UserID      uint64 `json:"user_id,omitempty"  form:"user_id,omitempty"`
}

type BookResponseDTO struct {
	Title string `json:"title" form:"title" binding:"required"`
	Year  string `json:"year" binding:"required"`
}

type BookUpdateDTO struct {
	ID          uint64 `json:"id" form:"id"`
	Title       string `json:"title" form:"title" binding:"required"`
	Year        string `json:"year" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	BookCover   string `json:"book_cover" form:"book_cover" `
	UserID      uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
}
