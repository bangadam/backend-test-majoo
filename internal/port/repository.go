package port

import (
	"context"

	"github.com/bangadam/backend-test-majoo/internal/domain"
	"github.com/bangadam/backend-test-majoo/pkg/paginator"
)

// Repository is the interface for the port layer.
type (
	UserRepository interface {
		GetUserByUsername(ctx context.Context, username string) (*domain.User, error)
	}

	MerchantRepository interface {
		GetMerchantByUserID(ctx context.Context, userID uint64) (*domain.Merchant, error)
	}

	OutletRepository interface {
		GetOutletByID(ctx context.Context, outletID uint64) (*domain.Outlet, error)
	}

	TransactionRepository interface {
		GetTransactionByMerchantID(ctx context.Context, merchantID uint64, req *domain.ReportTransactionRequest) ([]*domain.TransactionMerchant, *paginator.Pagination, error)
		GetTransactionByOutletID(ctx context.Context, outletID uint64, req *domain.ReportTransactionRequest) ([]*domain.TransactionOutlet, *paginator.Pagination, error)
	}
)