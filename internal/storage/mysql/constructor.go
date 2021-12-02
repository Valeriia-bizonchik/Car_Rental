package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type CarRentalStorage struct {
	db *gorm.DB
}

func NewCarRentalStorage(DNS string) (*CarRentalStorage, error) {
	db, err := gorm.Open("mysql", "mysql", DNS)
	if err != nil {
		return nil, errors.WithMessage(err, fmt.Sprintf(`failed to connect to the database DNS:%v`, DNS))
	}

	return &CarRentalStorage{db: db}, nil
}
