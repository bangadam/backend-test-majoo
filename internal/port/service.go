package port

import (
	"context"

	"github.com/bangadam/backend-test-majoo/internal/domain"
	"github.com/bangadam/backend-test-majoo/pkg/paginator"
)

// Service is the interface for the port layer.
type (
	AuthService interface {
		Login(ctx context.Context, req domain.LoginRequest) (*domain.LoginResponse, error)
	}

	ReportTransactionService interface {
		ReportTransactionMerchant(ctx context.Context, userID uint64, req domain.ReportTransactionRequest) (*domain.ReportMerchant, *paginator.Pagination, error)
		ReportTransactionOutlet(ctx context.Context, userID uint64, req domain.ReportTransactionRequest) (*domain.ReportOutlet, *paginator.Pagination, error)
	}
)