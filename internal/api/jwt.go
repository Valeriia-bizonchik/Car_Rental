package api

import (
	"github.com/Valeriia-bizonchik/CarRental/config"
	"net/http"
	"time"

	"github.com/Valeriia-bizonchik/CarRental/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func generateJWT(c *gin.Context, u models.User, timeLiveHour int) (string, error) {
	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(time.Duration(timeLiveHour) * time.Hour)
	// Create the JWT claims, which includes the username and expiry time
	claims := &models.Claims{
		UserID: u.ID,
		Role:   u.Role,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(config.JwtKey)
	if err != nil {
		return "", errors.Wrap(err, "failed to create JWT")
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	return tokenString, nil
}
