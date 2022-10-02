package domain

import (
	"time"

	"github.com/bangadam/backend-test-majoo/pkg/paginator"
)

type ReportMerchant struct {
	MerchantName string `json:"merchant_name"`
	Transactions []*TransactionMerchant `json:"transactions"`
}

type ReportTransactionOutletResponse struct {
	MerchantName string `json:"merchant_name"`
	OutletName string `json:"outlet_name"`
	Revenue float64 `json:"revenue"`
	Pagination *paginator.Pagination `json:"pagination"`
}

type ReportOutlet struct {
	MerchantName string `json:"merchant_name"`
	OutletName string `json:"outlet_name"`
	Transactions []*TransactionOutlet `json:"transactions"`
}

type ReportTransactionRequest struct {
	Pagination *paginator.Pagination `json:"pagination"`
	StartDate *time.Time `json:"start_date"`
	EndDate *time.Time `json:"end_date"`
	OutletID *uint64 `json:"outlet_id"`
}
