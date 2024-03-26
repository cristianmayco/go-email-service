package contract

import "time"

type EmailBlockDTO struct {
	Email        string    `json:"email" validate:"required,email"`
	CreationDate time.Time `validate:"required"`
}

func CreateEmailBlockDTO(email string) EmailBlockDTO {
	return EmailBlockDTO{
		Email:        email,
		CreationDate: time.Now(),
	}
}
