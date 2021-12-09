package postgres

import "github.com/Valeriia-bizonchik/CarRental/models"

func (c *CarRentalStorage) MigrateAllModels() error {
	err := c.DB.AutoMigrate(
		&models.User{},
		&models.Car{})
	if err != nil {
		return err
	}

	return nil
}
