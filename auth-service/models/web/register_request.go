package web

// RegisterRequest a object request for payload body register
type RegisterRequest struct {
	Fullname string `json:"fullname" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Address  string `json:"address" binding:"required"`
	City     string `json:"city" binding:"required"`
	Province string `json:"province" binding:"required"`
	Mobile   string `json:"mobile" binding:"required,max=12"`
	Password string `json:"password" binding:"required,min:3"`
	IsAdmin  bool   `json:"is_admin"`
}
