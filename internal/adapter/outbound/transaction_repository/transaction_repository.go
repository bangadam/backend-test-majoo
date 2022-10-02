package transaction_repository

import (
	"context"

	"github.com/bangadam/backend-test-majoo/internal/domain"
	"github.com/bangadam/backend-test-majoo/internal/port"
	"github.com/bangadam/backend-test-majoo/pkg/paginator"
	"gorm.io/gorm"
)

type transactionRepository struct {
	mysql *gorm.DB
}

func NewTransactionRepository(mysql *gorm.DB) port.TransactionRepository {
	return &transactionRepository{
		mysql: mysql,
	}
}

func (instance *transactionRepository) GetTransactionByMerchantID(ctx context.Context, merchantID uint64, req *domain.ReportTransactionRequest) ([]*domain.TransactionMerchant, *paginator.Pagination, error) {
	var (
		transactions []*domain.TransactionMerchant
		model 	  *domain.Transaction
		count int64
	)

	q := instance.mysql.Model(&model).
		Where("merchant_id = ?", merchantID).
		Preload("Merchant")
	// filter date
	startDate := req.StartDate
	endDate := req.EndDate
	q = q.Where("created_at BETWEEN ? AND ?", startDate, endDate).
		Select("DATE(created_at) as transaction_date, SUM(bill_total) as revenue").
		Group("DATE(created_at)").
		Count(&count).
		Order("DATE(created_at) ASC")
		
	
	// set pagination
	req.Pagination.Count = count
	req.Pagination = paginator.Paging(req.Pagination)
	err := q.Offset(req.Pagination.Offset).Limit(req.Pagination.Limit).Find(&transactions).Error
	if err != nil {
		return nil, nil, err
	}

	var response []*domain.TransactionMerchant
	for _, transaction := range transactions {
		response = append(response, transaction.TransactionMerchantTransformer())
	}

	return response, req.Pagination, nil
}

func (instance *transactionRepository) GetTransactionByOutletID(ctx context.Context, outletID uint64, req *domain.ReportTransactionRequest) ([]*domain.TransactionOutlet, *paginator.Pagination, error) {
	var (
		transactions []*domain.TransactionOutlet
		model 	  *domain.Transaction
		count int64
	)

	q := instance.mysql.Model(&model).
		Where("outlet_id = ?", outletID).
		Preload("Outlet")
	// filter date
	startDate := req.StartDate
	endDate := req.EndDate
	q = q.Where("created_at BETWEEN ? AND ?", startDate, endDate).
		Select("DATE(created_at) as transaction_date, SUM(bill_total) as revenue").
		Group("DATE(created_at)").
		Count(&count).
		Order("DATE(created_at) ASC")
		
	
	// set pagination
	req.Pagination.Count = count
	req.Pagination = paginator.Paging(req.Pagination)
	err := q.Offset(req.Pagination.Offset).Limit(req.Pagination.Limit).Find(&transactions).Error
	if err != nil {
		return nil, nil, err
	}

	var response []*domain.TransactionOutlet
	for _, transaction := range transactions {
		response = append(response, transaction.TransactionOutletTransformer())
	}

	return response, req.Pagination, nil
}