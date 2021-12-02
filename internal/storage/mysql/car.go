package mysql

import "github.com/Valeriia-bizonchik/CarRental/models"

/*

func (c *Car) CreateCar() *Car {
	db.NewRecord(c)
	db.Create(&c)
	return c

}

func GetAllCars() []Car {
	var Cars []Car
	db.Find(&Cars)
	return Cars

}

func GetCarById(Id int64) (*Car, *gorm.DB) {
	var getCar Car
	db := db.Where("ID=?", Id).Find(&getCar)
	return &getCar, db

}

func DeleteCar(ID int64) Car {
	var car Car
	db.Where("ID=?", ID).Delete(car)
	return car

}
*/

func (s *CarRentalStorage) CreateCar(car models.Car) (*models.Car, error) {
	car.ID = 0
	err := s.db.Create(&car).Error
	if err != nil {
		return nil, err
	}

	return &car, nil
}

func (s *CarRentalStorage) GetAllCars() ([]*models.Car, error) {
	var cars []*models.Car
	err := s.db.Find(&cars).Error
	if err != nil {
		return nil, err
	}
	return cars, nil
}
