package domain

import "time"

type Outlet struct {
	ID uint64 `gorm:"primary_key;column:id"`
	CreatedBy uint64 `gorm:"column:created_by"`
	UpdatedBy uint64 `gorm:"column:updated_by"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`

	MerchantID uint64 `gorm:"column:merchant_id"`
	OutletName string `gorm:"column:outlet_name"`
}

type PublicOutlet struct {
	OutletName string `json:"outlet_name" gorm:"column:outlet_name"`
}

func (u *Outlet) IsEmpty() bool {
	return u == nil
}

func (u *Outlet) IsMyOutlet(MerchantID uint64) bool {
	return u.MerchantID == MerchantID
}