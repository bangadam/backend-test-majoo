package domain

import "time"

type Transaction struct {
	ID uint64 `gorm:"primary_key;column:id"`
	CreatedBy uint64 `gorm:"column:created_by"`
	UpdatedBy uint64 `gorm:"column:updated_by"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	
	MerchantID uint64 `gorm:"column:merchant_id"`
	OutletID uint64 `gorm:"column:outlet_id"`
	BillTotal float64 `gorm:"column:bill_total"`

	Outlet Outlet `gorm:"foreignkey:outlet_id"`
}

type TransactionMerchant struct {
	TransactionDate string  `json:"transaction_date" gorm:"column:transaction_date"`
	Revenue         float64 `json:"revenue" gorm:"column:revenue"`
	Merchant *Merchant `json:"merchant,omitempty"`
}

// ToTrasformer
func (t *TransactionMerchant) TransactionMerchantTransformer() *TransactionMerchant {
	// format date string to 2019-01-01
	t.TransactionDate = t.TransactionDate[:10]
	return &TransactionMerchant{
		TransactionDate: t.TransactionDate,
		Revenue:         t.Revenue,
	}
}


type TransactionOutlet struct {
	TransactionDate string  `json:"transaction_date" gorm:"column:transaction_date"`
	Revenue         float64 `json:"revenue" gorm:"column:revenue"`
	Outlet *Outlet `json:"outlet,omitempty"`
}

// ToTransformer
func (t *TransactionOutlet) TransactionOutletTransformer() *TransactionOutlet {
	// format date string to 2019-01-01
	t.TransactionDate = t.TransactionDate[:10]
	return &TransactionOutlet{
		TransactionDate: t.TransactionDate,
		Revenue:         t.Revenue,
	}
}