package outlet_repository

import (
	"context"
	"errors"

	"github.com/bangadam/backend-test-majoo/internal/domain"
	"github.com/bangadam/backend-test-majoo/internal/port"
	"gorm.io/gorm"
)

type outletRepository struct {
	mysql *gorm.DB
}

func NewOutletRepository(mysql *gorm.DB) port.OutletRepository {
	return &outletRepository{mysql: mysql}
}

func (instance *outletRepository) GetOutletByID(ctx context.Context, outletID uint64) (*domain.Outlet, error) {
	var (
		outlet domain.Outlet
	)

	err := instance.mysql.Where("id = ?", outletID).First(&outlet).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	
	return &outlet, nil
}