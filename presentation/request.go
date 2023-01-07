package presentation

type AuthRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AddBookRequestBody struct {
	Title  string
	Author string
}

type EditBookRequestBody struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type DeleteBookRequestBody struct {
	ID int `json:"id"`
}

type BorrowBookRequestBody struct {
	BookID int `json:"book_id"`
}
