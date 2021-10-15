package model

type (
	User struct {
		BaseModel
		Email    string `json:"email" gorm:"column:email"`
		Password string `json:"password" gorm:"column:password"`
	}

	UserRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	ChangePassword struct {
		Password    string `json:"password"`
		NewPassword string `json:"new-password"`
	}

	DeleteRequest struct {
		Password string `json:"password"`
	}
)
