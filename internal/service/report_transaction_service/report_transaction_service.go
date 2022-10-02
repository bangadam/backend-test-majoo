package report_transaction_service

import (
	"context"
	"time"

	"github.com/bangadam/backend-test-majoo/internal/domain"
	"github.com/bangadam/backend-test-majoo/internal/port"
	"github.com/bangadam/backend-test-majoo/pkg/paginator"
	"github.com/bangadam/backend-test-majoo/pkg/response"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type reportTransactionService struct {
	log *zap.Logger
	merchantRepository port.MerchantRepository
	outletRepository port.OutletRepository
	transactionRepository port.TransactionRepository
}

func NewReportTransactionService(log *zap.Logger, merchantRepository port.MerchantRepository, outletRepository port.OutletRepository, transactionRepository port.TransactionRepository) port.ReportTransactionService {
	return &reportTransactionService{
		log: log,
		merchantRepository: merchantRepository,
		outletRepository: outletRepository,
		transactionRepository: transactionRepository,
	}
}

func (instance *reportTransactionService) ReportTransactionMerchant(ctx context.Context, userID uint64, req domain.ReportTransactionRequest) (*domain.ReportMerchant, *paginator.Pagination, error) {
	// get merchant
	merchant, err := instance.merchantRepository.GetMerchantByUserID(ctx, userID)
	if err != nil {
		instance.log.Error("failed to get merchant", zap.Error(err))
		return nil, nil, err
	}

	if merchant.IsEmpty() {
		instance.log.Error("merchant not found", zap.Uint64("user_id", userID))
		return nil, nil, nil
	}

	// check has filter start date and end date
	if req.StartDate == nil && req.EndDate == nil {
		// set default date 1 November 2021 to 30 November 2021
		startDate := time.Date(2021, 11, 1, 0, 0, 0, 0, time.Local)
		req.StartDate = &startDate
		endDate := time.Date(2021, 11, 30, 0, 0, 0, 0, time.Local)
		req.EndDate = &endDate
	}

	// get transaction
	transactions, pagination, err := instance.transactionRepository.GetTransactionByMerchantID(ctx, merchant.ID, &req)
	if err != nil {
		instance.log.Error("failed to get transaction", zap.Error(err))
		return nil, nil, err
	}

	// transform
	response := &domain.ReportMerchant{
		MerchantName: merchant.MerchantName,
		Transactions: transactions,
	}

	return response, pagination, nil
}

func (instance *reportTransactionService) ReportTransactionOutlet(ctx context.Context, userID uint64, req domain.ReportTransactionRequest) (*domain.ReportOutlet, *paginator.Pagination, error) {
	// get merchant
	merchant, err := instance.merchantRepository.GetMerchantByUserID(ctx, userID)
	if err != nil {
		instance.log.Error("failed to get merchant", zap.Error(err))
		return nil, nil, response.New(fiber.StatusInternalServerError, response.WithMessage("failed to get merchant"))
	}

	if merchant.IsEmpty() {
		instance.log.Error("merchant not found", zap.Uint64("user_id", userID))
		return nil, nil, response.New(fiber.StatusNotFound, response.WithMessage("merchant not found"))
	}

	if req.OutletID == nil {
		instance.log.Error("outlet id not found", zap.Uint64("user_id", userID))
		return nil, nil, response.New(fiber.StatusNotFound, response.WithMessage("outlet id not found"))
	}

	outlet, err := instance.outletRepository.GetOutletByID(ctx, *req.OutletID)

	if err != nil {
		instance.log.Error("failed to get outlet", zap.Error(err))
		return nil, nil, response.New(fiber.StatusInternalServerError, response.WithMessage("failed to get outlet"))
	}

	if outlet.IsEmpty() {
		instance.log.Error("outlet not found", zap.Uint64("outlet_id", *req.OutletID))
		return nil, nil, response.New(fiber.StatusNotFound, response.WithMessage("outlet not found"))
	}

	if !outlet.IsMyOutlet(merchant.ID) {
		instance.log.Error("you are not allowed to access this outlet", zap.Uint64("outlet_id", *req.OutletID))
		return nil, nil, response.New(fiber.StatusUnauthorized, response.WithMessage("you are not allowed to access this outlet"))
	}

	// check has filter start date and end date
	if req.StartDate == nil && req.EndDate == nil {
		// set default date 1 November 2021 to 30 November 2021
		startDate := time.Date(2021, 11, 1, 0, 0, 0, 0, time.Local)
		req.StartDate = &startDate
		endDate := time.Date(2021, 11, 30, 0, 0, 0, 0, time.Local)
		req.EndDate = &endDate
	}

	// get transaction
	transactions, pagination, err := instance.transactionRepository.GetTransactionByOutletID(ctx, outlet.ID, &req)
	if err != nil {
		instance.log.Error("failed to get transaction", zap.Error(err))
		return nil, nil, err
	}

	// transform
	response := &domain.ReportOutlet{
		MerchantName: merchant.MerchantName,
		OutletName: outlet.OutletName,
		Transactions: transactions,
	}

	return response, pagination, nil
}