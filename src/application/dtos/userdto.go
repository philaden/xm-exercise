package dtos

type (
	LoginDto struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	TokenDto struct {
		Email  string `json:"email"`
		Token  string `json:"token"`
		UserId string `json:"userId"`
	}

	CreateUserDto struct {
		FirstName   string `json:"firstName"`
		LastName    string `json:"lastName"`
		PhoneNumber string `json:"phoneNumber"`
		Email       string `json:"email"`
		Password    string `json:",omitempty"`
		Role        string `json:"role" default:"user"`
	}

	UserResponseDto struct {
		ID          string `json:"id"`
		FirstName   string `json:"firstName"`
		LastName    string `json:"lastName"`
		PhoneNumber string `json:"phoneNumber"`
		Email       string `json:"email"`
	}

	UsersResponseDto []UserResponseDto
)
