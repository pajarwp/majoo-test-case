package user

type UserLoginModel struct {
	Username string `json:"user_name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserLoginDataModel struct {
	ID       int    `json:"id"`
	Password string `json:"password"`
}
