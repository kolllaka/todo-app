package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/KoLLlaka/todo-app/internal/todo"
	"github.com/KoLLlaka/todo-app/pkg/repository"

	"github.com/dgrijalva/jwt-go"
)

const (
	salt = "hjhk213haejhw"

	jwtTokenTTL   = 12 * time.Hour
	jwtSigningKey = "qweqdq123asdaf"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func newAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)

	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	// get user from db
	userID, err := s.repo.GetUserID(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(jwtTokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userID.Id,
	})

	return token.SignedString([]byte(jwtSigningKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(jwtSigningKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserID, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
