package services

import (
	"errors"
	"os"
	"github.com/cesc1802/go_training/internal/dto/requests"
	"github.com/cesc1802/go_training/internal/dto/responses"
	"github.com/cesc1802/go_training/internal/storages"
	"golang.org/x/crypto/bcrypt"
)

func Login(loginPayload dto_request.LoginPayload) (dto_response.Login, error) {
	user := &storages.User{}
	storages.Get().Take(user)
	if user == (&storages.User{}) {
		return dto_response.Login{}, errors.New("user not found")
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginPayload.Password))
	if err != nil {
		return dto_response.Login{}, errors.New("wrong password")
	}

	jwtWrapper := JwtWrapper{
		SecretKey:       os.Getenv("SECRET"),
		Issuer:          "AuthService",
		ExpirationHours: 9999,
	}

	bearToken, err := jwtWrapper.GenerateToken(loginPayload.UserId)
	if err != nil {
		return dto_response.Login{}, errors.New("cannot generate token")
	}

	return dto_response.Login{Token: bearToken}, err
}