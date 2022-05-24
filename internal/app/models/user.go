package models

//User model defeniton
type User struct {
	Id       int    `json:"id_user"`
	Login    string `json:"login_user"`
	Password string `json:"password_user"`
	Role     string `json:"role"`
}
