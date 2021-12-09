package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CarRentalStorage struct {
	*gorm.DB
}

func NewCarRentalStorage(DNS string) (*CarRentalStorage, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  DNS,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return &CarRentalStorage{DB: db}, nil
}
