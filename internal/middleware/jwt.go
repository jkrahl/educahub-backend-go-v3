package middleware

import (
	"context"
	"educahub/configs"
	"log"
	"net/http"
	"net/url"
	"time"

	jwtutils "educahub/internal/jwt"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	jwks "github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
	adapter "github.com/gwatts/gin-adapter"
)

type CustomClaims struct {
	Scope string `json:"scope"`
}

func (c CustomClaims) Validate(ctx context.Context) error {
	return nil
}

func GetAuthMiddleware() (gin.HandlerFunc, error) {
	issuerURL, err := url.Parse("https://" + configs.GetViperString("AUTH0_DOMAIN") + "/")
	if err != nil {
		return nil, err
	}
	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{configs.GetViperString("AUTH0_AUDIENCE")},
		validator.WithCustomClaims(
			func() validator.CustomClaims {
				return &CustomClaims{}
			},
		),
		validator.WithAllowedClockSkew(time.Minute),
	)
	if err != nil {
		return nil, err
	}

	errorHandler := func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("Encountered error while validating JWT: %v", err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message":"Failed to validate JWT."}`))
	}

	middleware := jwtmiddleware.New(
		jwtValidator.ValidateToken,
		jwtmiddleware.WithErrorHandler(errorHandler),
	)

	return adapter.Wrap(middleware.CheckJWT), nil
}

func SetSubInContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		sub, err := jwtutils.GetSubFromTokenFromContext(c)
		if err != nil {
			c.JSON(401, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		c.Set("sub", sub)
		c.Next()
	}
}
