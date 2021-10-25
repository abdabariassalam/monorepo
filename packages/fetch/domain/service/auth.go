package service

import (
	"errors"

	"github.com/bariasabda/monorepo/packages/fetch/constants"
	"github.com/golang-jwt/jwt"
)

type User struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

func (s *service) VerifyToken(reqToken string) (*User, error) {
	user := &User{}
	_, err := jwt.ParseWithClaims(reqToken, user, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.cfg.JWT.SecretKey), nil
	})
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, constants.ErrExpiredToken) {
			return nil, constants.ErrExpiredToken
		}
		return nil, constants.ErrInvalidToken
	}
	return user, nil

}
