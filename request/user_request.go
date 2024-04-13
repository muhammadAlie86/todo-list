package request

type Register struct {
	Email                string `form:"email" validate:"required,email" example:"todotest@gmail.com"`
	Username             string `form:"username" validate:"required,min=3,max=16" example:"todo"`
	FullName             string `form:"sn" validate:"required,min=3" example:"Todo"`
	RoleId               int16  `form:"role_id" validate:"required" example:"1"`
	Password             string `form:"password" validate:"min=8" example:"12345678"`
	ConfirmationPassword string `form:"confirmation_password" validate:"min=8" example:"12345678"`
}

type Login struct {
	Username string `json:"username" validate:"omitempty,min=3,max=50" example:"todo"`
	Password string `json:"password" validate:"min=8" example:"12345678"`
}

type Refresh struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type Verification struct {
	Username         string `form:"verification_username" validate:"required,min=3,max=16" example:"todo"`
	Email            string `form:"email" validate:"required,email" example:"todotest@gmail.com"`
	VerificationCode string `form:"verification_code" validate:"required,min=6,max=6"`
}
