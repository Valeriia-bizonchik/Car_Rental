package api

import (
	"github.com/Valeriia-bizonchik/CarRental/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (a *API) UserInfo(c *gin.Context) {
	claims, isExists := c.Get("user_info")
	if !isExists {
		c.JSON(http.StatusInternalServerError, "failed get data about authorised user")
		return
	}

	clm := claims.(*models.Claims)

	c.JSON(http.StatusOK, clm)
	return
}
