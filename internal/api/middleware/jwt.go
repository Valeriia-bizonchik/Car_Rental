package middleware

import (
	"github.com/Valeriia-bizonchik/CarRental/config"
	"github.com/Valeriia-bizonchik/CarRental/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ValidateJWT(c *gin.Context) {
	// We can obtain the session token from the requests cookies, which come with every request
	tknStr, err := c.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// For any other type of error, return a bad request status
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// Initialize a new instance of `Claims`
	claims := &models.Claims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return config.JwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Set("user_info", claims)

	c.Next()
}
