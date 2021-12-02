package routes

import (
	"github.com/Valeriia-bizonchik/CarRental/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterCarShareRoutes = func(router *mux.Router) {
	router.HandleFunc("/car/", controllers.CreateCar).Methods("POST")
	router.HandleFunc("/car/", controllers.GetCar).Methods("GET")
	router.HandleFunc("/car/{carId}", controllers.GetCarById).Methods("GET")
	router.HandleFunc("/car/{carId}", controllers.UpdateCar).Methods("PUT")
	router.HandleFunc("/car/{carId}", controllers.DeleteCar).Methods("DELETE")
}
