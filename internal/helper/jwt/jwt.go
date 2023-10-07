package jwt

import (
	"ecommercestore/internal/helper"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
	"time"
)

var (
	ErrInvalidToken = helper.NewCustomError(
		errors.New("invalid token"),
		"invalid token",
		"ErrInvalidToken",
		http.StatusUnauthorized,
	)
	ErrGenerateToken = helper.NewCustomError(
		errors.New("cannot generate token"),
		"cannot generate token",
		"ErrGenerateToken",
		http.StatusInternalServerError,
	)
)

type jwtProvider struct {
	secret string
}

func NewJWTProvider(secret string) *jwtProvider {
	return &jwtProvider{secret: secret}
}

func (j *jwtProvider) GenerateToken(id int, expiry int) string {
	nowTime := time.Now().UTC()
	expireTime := nowTime.Add(time.Second * time.Duration(expiry))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			ID:        fmt.Sprint(id),
		})
	tokenString, err := token.SignedString([]byte(j.secret))
	if err != nil {
		log.Panic().Err(err).Msg("Error building JWT")
	}
	return tokenString
}

func (j *jwtProvider) ValidateToken(myToken string) (int, error) {
	token, err := jwt.ParseWithClaims(myToken, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.secret), nil
	})
	if err != nil {
		log.Error().Err(err).Str("tokenStr", myToken).Msg("Error parsing JWT")
		return 0, ErrInvalidToken
	}

	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		id, err := strconv.Atoi(claims.ID)
		if err != nil {
			log.Error().Err(err).Msg("Error unmarshalling JWT claims")
			return 0, ErrInvalidToken
		}

		return id, nil
	}

	return 0, ErrInvalidToken
}
