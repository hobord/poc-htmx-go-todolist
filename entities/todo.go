package entities

type Todo struct {
	ID          string `json:"id"`
	Group       string `json:"group"`
	Priority    int    `json:"priority"`
	User        string `json:"user"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}
