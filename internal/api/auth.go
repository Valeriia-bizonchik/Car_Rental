package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Valeriia-bizonchik/CarRental/config"
	"github.com/Valeriia-bizonchik/CarRental/models"
	"github.com/Valeriia-bizonchik/CarRental/utils"
	"github.com/gin-gonic/gin"
)

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (a *API) Login(c *gin.Context) {
	var loginParams LoginReq
	err := c.BindJSON(&loginParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Println(config.PrettyPrint(loginParams))

	var u models.User
	u.Name = loginParams.Username
	err = a.storage.Where(&u).Find(&u).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if u.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid username or password",
		})
		return
	}

	if !utils.CheckPasswordHash(loginParams.Password, u.Passwd) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid username or password",
		})
		return
	}

	token, err := generateJWT(c, u, 24)
	if err != nil {
		a.log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"token":  token,
		"user":   u,
	})
	return
}

type RegisterReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (a *API) Register(c *gin.Context) {
	var registerParams RegisterReq
	err := c.BindJSON(&registerParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	config.PrettyPrint(registerParams)

	var u models.User
	u.Name = registerParams.Username
	u.Passwd = registerParams.Password
	u.Role = models.Customer

	err = a.storage.DB.Create(&u).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := generateJWT(c, u, 24)
	if err != nil {
		a.log.Error(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"token":  token,
		"user":   u,
	})
}

func (a *API) Logout(c *gin.Context) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Unix(0, 0),
	})

	c.JSON(http.StatusOK, gin.H{
		"status": "success logout",
	})
}
