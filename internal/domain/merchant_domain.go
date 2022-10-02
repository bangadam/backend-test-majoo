package domain

import (
	"time"
)

type Merchant struct {
	ID uint64 `gorm:"primary_key;column:id"`
	UserID uint64 `gorm:"column:user_id"`
	MerchantName string `gorm:"column:merchant_name"`
	CreatedBy uint64 `gorm:"column:created_by"`
	UpdatedBy uint64 `gorm:"column:updated_by"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

// check if empty
func (instance *Merchant) IsEmpty() bool {
	return instance == nil
}


type PublicMerchant struct {
	MerchantName string `json:"merchant_name"`
}