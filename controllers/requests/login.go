package requests

type RegisterRequest struct {
	UserName   string  `form:"username"`
	Password   string  `form:"password"`
	RePassword string  `form:"repassword"`
	Email      string  `form:"email"`
}
