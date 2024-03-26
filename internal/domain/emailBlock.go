package domain

import "time"

type EmailBlock struct {
	Email         string `validate:"required,email"`
	AbstractModel `validate:"required"`
}

func CreateEmailBlock(email string) EmailBlock {
	return EmailBlock{
		Email: email,
		AbstractModel: AbstractModel{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

}
