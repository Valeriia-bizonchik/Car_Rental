package storage

import "github.com/Valeriia-bizonchik/CarRental/models"

type CarRental interface {
	CreateCar(car models.Car) (*models.Car, error)
	GetAllCars() ([]*models.Car, error)
}
