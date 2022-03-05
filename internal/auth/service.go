package auth

import (
	"context"
	"started_kit/internal/entities"
	"started_kit/internal/errors"
	"started_kit/pkg/log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	Login(ctx context.Context, email, password string) (string, error)
}

type Identity interface {
	GetID() string
}

type service struct {
	accessToken  string
	accessTTL    int
	refreshToken string
	refreshTTL   int
	logger       log.Logger
}

func NewService(accessToken string, accessTTL int, refreshToken string, refreshTTL int, logger log.Logger) Service {
	return service{accessToken, accessTTL, refreshToken, refreshTTL, logger}
}
func (s service) Login(ctx context.Context, email, password string) (string, error) {
	if identity := s.authenticate(ctx, email, password); identity != nil {
		return s.generateJWT(identity)
	}
	return "", errors.Unauthorized("")
}

func (s service) authenticate(ctx context.Context, email, password string) Identity {
	logger := s.logger.With(ctx, "users", email)
	// TODO: the following authentication logic is only for demo purpose
	if email == "demo" && password == "pass" {
		logger.Infof("authentication successful")
		return entities.User{}
	}

	logger.Infof("authentication failed")
	return nil
}

func (s service) generateJWT(Identity Identity) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  Identity.GetID(),
		"exp": time.Now().Add(time.Duration(s.accessTTL) * time.Hour).Unix(),
	}).SignedString([]byte(s.accessToken))
}
