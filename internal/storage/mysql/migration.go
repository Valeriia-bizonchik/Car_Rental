package mysql

import "github.com/Valeriia-bizonchik/CarRental/models"

func (s *CarRentalStorage) MigrateAllModels() {
	s.db.AutoMigrate(&models.Car{})
}
