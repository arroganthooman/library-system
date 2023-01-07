package presentation

type ResponseStandard struct {
	Status          int         `json:"status"`
	ErrorMessage    string      `json:"error_message"`
	RawErrorMessage string      `json:"raw_error_message"`
	Data            interface{} `json:"data"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type BookResponse struct {
	Author     string `json:author`
	Title      string `json:title`
	IsBorrowed bool   `json:"is_borrowed"`
}

type AllBookResponse struct {
	Author     string `json:author`
	Title      string `json:title`
	IsBorrowed bool   `json:"is_borrowed"`
	ID         int    `json:"id"`
}
