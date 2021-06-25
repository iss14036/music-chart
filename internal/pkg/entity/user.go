package entity

import "errors"

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
	Hobby    string `json:"hobby"`
	Gender   string `json:"gender"`
	Address  string `json:"address"`
}

type UserFilter struct {
	ID       string
	Username string
	Password string
}

type UserToken struct {
	Token string `json:"token"`
}

var (
	ErrorEmptyField = errors.New("There is field that empty. Please check again")
)

func (u *User) Validate() error {
	if u.Username == "" || u.Password == "" || u.FullName == "" || u.Address == "" || u.Gender == "" || u.Hobby == "" {
		return ErrorEmptyField
	}

	return nil
}
