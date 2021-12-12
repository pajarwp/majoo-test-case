package usecase

import (
	"encoding/hex"
	"errors"
	"majoo-test-case/config"
	"majoo-test-case/entity/user"
	"majoo-test-case/entity/user/repository"
	"time"

	"crypto/md5"

	"github.com/dgrijalva/jwt-go"
)

type UserUsecase interface {
	UserLogin(um *user.UserLoginModel) (string, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(u repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: u,
	}
}

func (u *userUsecase) UserLogin(um *user.UserLoginModel) (string, error) {
	userObj, err := u.userRepository.UserLogin(um)
	if err != nil {
		return "", err
	}
	err = verifyPassword(um.Password, userObj.Password)
	if err != nil {
		return "", err
	}
	token, err := createJWTToken(userObj)
	if err != nil {
		return "", err
	}
	return token, nil
}

func verifyPassword(password string, hashedpwd string) error {
	hasher := md5.New()
	hasher.Write([]byte(password))
	password = hex.EncodeToString(hasher.Sum(nil))

	if password != hashedpwd {
		return errors.New("Login Failed, Incorrect Password")
	}

	return nil
}

func createJWTToken(u *user.UserLoginDataModel) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = u.ID
	claims["exp"] = time.Now().Add(time.Hour * 8).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.JWT_SECRET))
	if err != nil {
		return "", err
	}
	return t, nil
}
