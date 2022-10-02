package merchant_repository

import (
	"context"
	"errors"

	"github.com/bangadam/backend-test-majoo/internal/domain"
	"github.com/bangadam/backend-test-majoo/internal/port"
	"gorm.io/gorm"
)

type merchantRepository struct {
	mysql *gorm.DB
}

func NewMerchantRepository(mysql *gorm.DB) port.MerchantRepository {
	return &merchantRepository{mysql: mysql}
}

func (instance *merchantRepository) GetMerchantByUserID(ctx context.Context, userID uint64) (*domain.Merchant, error) {
	var (
		merchant domain.Merchant
	)

	err := instance.mysql.Where("user_id = ?", userID).First(&merchant).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	
	return &merchant, nil
}