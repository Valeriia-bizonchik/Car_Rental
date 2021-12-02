package v1

import (
	"github.com/Valeriia-bizonchik/CarRental/internal/api/rest/middleware"
	"github.com/Valeriia-bizonchik/CarRental/internal/storage"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type API struct {
	r *gin.Engine

	storage storage.CarRental
	log     *zap.SugaredLogger
}

func NewAPI(storage storage.CarRental, log *zap.SugaredLogger) *API {
	return &API{storage: storage, log: log}

}

func (a *API) InitRoutes() {
	r := gin.Default()

	// Setup middleware
	r.Use(middleware.ContentType)

	// status
	r.GET("/status", func(context *gin.Context) {
		context.JSON(200, gin.H{"status": "ok"})
	})

	v1 := r.Group(`/api/v1`)
	{
		v1.POST(`/echo`, a.Echo)
		//v1.GET("/cars/get_all", a.GetAllCars)

		car := v1.Group("/car")
		{
			//car.POST("/create", a.CreteCar)
			//car.GET("/get", a.GetCar)
			car.GET("/get_all", a.GetAllCars)
			/*
				router.HandleFunc("/car/", controllers.CreateCar).Methods("POST")
				router.HandleFunc("/car/", controllers.GetCar).Methods("GET")
				router.HandleFunc("/car/{carId}", controllers.GetCarById).Methods("GET")
				router.HandleFunc("/car/{carId}", controllers.UpdateCar).Methods("PUT")
				router.HandleFunc("/car/{carId}", controllers.DeleteCar).Methods("DELETE")
			*/
		}
	}

	a.r = r
}

func (a *API) Run(host, port string) error {
	return a.r.Run(host + ":" + port)
}
