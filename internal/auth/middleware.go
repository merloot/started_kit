package auth

import (
	"context"
	"started_kit/internal/entities"

	"github.com/dgrijalva/jwt-go"
	routing "github.com/go-ozzo/ozzo-routing/v2"
)

// Handler returns a JWT-based authentication middleware.
func Handler(verificationKey string) routing.Handler {
	return nil
	// return auth.JWT(verificationKey, auth.JWTOptions{TokenHandler: handleToken})
}

// handleToken stores the user identity in the request context so that it can be accessed elsewhere.
func handleToken(c *routing.Context, token *jwt.Token) error {
	ctx := WithUser(
		c.Request.Context(),
		token.Claims.(jwt.MapClaims)["id"].(string),
	)
	c.Request = c.Request.WithContext(ctx)
	return nil
}

type contextKey int

const (
	userKey contextKey = iota
)

func WithUser(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, userKey, entities.User{})
}

func CurrentUser(ctx context.Context) Identity {
	if user, ok := ctx.Value(userKey).(entities.User); ok {
		return user
	}
	return nil
}
