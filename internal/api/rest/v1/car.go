package v1

import (
	"github.com/Valeriia-bizonchik/CarRental/models"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func (a *API) Echo(c *gin.Context) {
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		err = c.AbortWithError(400, err)
		a.log.Error(err)
		return
	}
	_, err = c.Writer.Write(data)
	if err != nil {
		a.log.Error(err)
		c.AbortWithStatus(500)
	}
	c.Status(200)
	return
}

func (a *API) GetAllCars(c *gin.Context) {
	//cars, err := a.storage.GetAllCars()
	//if err != nil {
	//	a.log.Error(err)
	//	c.AbortWithError(500, errors.New("internal server error"))
	//}

	cars := []*models.Car{
		{
			Name:        "name_1",
			Author:      "author_1",
			Publication: "publication_1",
		},
		{
			Name:        "name_2",
			Author:      "author_2",
			Publication: "publication_2",
		},
		{
			Name:        "name_3",
			Author:      "author_3",
			Publication: "publication_3",
		},
	}

	c.JSON(200, cars)
	return
}
