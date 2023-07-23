package entity

type Book struct {
	ID          uint64 `json:"id"`
	Title       string `json:"title" binding:"required"`
	Year        string `json:"year"`
	Description string `json:"description"`
	BookCover   string `json:"book_cover"`
	UserID      uint64 `json:"-"`
	User        User   `json:"user"`
}
