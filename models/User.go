package models

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) Validate() *Error {
	var error Error

	if u.Email == "" {
		error.Message = "Email is missing"
		return &error
	}

	if u.Password == "" {
		error.Message = "Password is missing"
		return &error
	}

	return nil
}
