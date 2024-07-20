package models

type User struct {
	ID       int64  `json:"id" validate:"omitempty"`
	Username string `json:"username" validate:"required,alphanumunicode,min=3,max=255"`
	Email    string `json:"email" validate:"required,email,max=255"`
	Age      uint8  `json:"age" validate:"min=0,max=255"`
}

type UsersList struct {
	TotalCount int     `json:"total_count"`
	Users      []*User `json:"users"`
}
