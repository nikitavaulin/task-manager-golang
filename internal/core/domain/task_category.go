package domain

type TaskCategory struct {
	ID           int64  `json:"id"`
	CategoryName string `json:"category_name"`

	UserID int64 `json:"user_id"`
}
