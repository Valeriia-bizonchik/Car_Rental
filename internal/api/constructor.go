package api

import (
	"github.com/Valeriia-bizonchik/CarRental/internal/api/middleware"
	"github.com/Valeriia-bizonchik/CarRental/internal/storage/postgres"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type API struct {
	r *gin.Engine

	storage *postgres.CarRentalStorage
	log     *zap.SugaredLogger
}

func NewAPI(storage *postgres.CarRentalStorage, log *zap.SugaredLogger) *API {
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

	api := r.Group(`/api`)
	{
		api.POST(`/echo`, a.Echo)

		// handle auth
		api.POST("/login", a.Login)
		api.POST("/register", a.Register)
		api.GET("/logout", a.Logout)

		car := api.Group("/car")
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

		secret := api.Group("/only_auth", middleware.ValidateJWT)
		{
			secret.GET("/user_info", a.UserInfo)
		}

		adminOnly := api.Group("/admin", middleware.ValidateAdmin)
		{
			adminOnly.GET("/user_info")
		}
	}

	a.r = r
}

func (a *API) Run(host, port string) error {
	return a.r.Run(host + ":" + port)
}
