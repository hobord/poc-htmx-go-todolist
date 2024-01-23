package entities

type UserEntity interface {
	GetUserID() string
}

type TodoGroup struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	Title  string `json:"title"`
	Color  string `json:"color,omitempty"`
	Todos  []*TodoItem
}

func (t *TodoGroup) GetUserID() string {
	return t.UserID
}

type TodoItem struct {
	ID      string `json:"id"`
	GroupID string `json:"group_id"`
	UserID  string `json:"user_id"`
	Title   string `json:"title"`
	// Description string `json:"description"`
	Completed bool `json:"completed"`
	Priority  int  `json:"priority,omitempty"`
}

func (t *TodoItem) GetUserID() string {
	return t.UserID
}
