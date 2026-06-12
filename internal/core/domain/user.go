package domain

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
}
