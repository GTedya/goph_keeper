package token

import (
	"errors"
)

var (
	// ErrExpiredToken закончился срок действия токена
	ErrExpiredToken = errors.New("token has expired")
)

// Manager интерфейс генерации и проверки токенов для аутентификации
type Manager interface {
	// Create создает токен для указанного userID
	Create(userID int) (token string, err error)
	// Validate проверяет токен на валидность
	Validate(accessToken string) (payload *Payload, err error)
}
