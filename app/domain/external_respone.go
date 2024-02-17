package domain

type ExternalResponse struct {
	UserId int    `json:"userId" validate:"required"`
	ID     int    `json:"id" validate:"required"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
