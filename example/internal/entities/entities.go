package entities

type Todo struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedBy   string `json:"createdBy"`
}
