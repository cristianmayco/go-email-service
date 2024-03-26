package main

import (
	"github.com/cristianmayco/go-email-service/internal/domain"
)

func main() {
	email := "teste@gmail.com"

	emailBlock1 := domain.CreateEmailBlock(email)

	println(emailBlock1.CreatedAt.Day())

}
