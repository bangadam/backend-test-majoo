package report_handler

import (
	"strconv"

	"github.com/bangadam/backend-test-majoo/internal/domain"
	"github.com/bangadam/backend-test-majoo/internal/middleware"
	"github.com/bangadam/backend-test-majoo/internal/port"
	"github.com/bangadam/backend-test-majoo/pkg/paginator"
	"github.com/bangadam/backend-test-majoo/pkg/response"
	"github.com/gofiber/fiber/v2"
)

type reportHandler struct {
	app *fiber.App
	reportTransactionService port.ReportTransactionService
}

func NewReportHandler(app *fiber.App, reportTransactionService port.ReportTransactionService)  {
	reportHandler := &reportHandler{
		app: app,
		reportTransactionService: reportTransactionService,
	}
	// set middleware
	api := reportHandler.app.Group("/api/v1")
	api.Get("/report-transaction-merchant", middleware.Authenticate(), reportHandler.DoReportTransactionMerchant)
	api.Get("/report-transaction-outlet/:outlet_id", middleware.Authenticate(), reportHandler.DoReportTransactionOutlet)
}

// Report Transaction Merchant godoc
// @Summary      Report Transaction Merchant
// @Description  API for get report transaction merchant
// @Tags         Report
// @Accept       json
// @Produce      json
// @Param        page   query     int  false  "page filter"  	   @Default("1")
// @Param        limit  query     int  false  "data limit filter"  @Default("10")
// @Param        start_date  query     string  false  "start date filter"
// @Param        end_date  query     string  false  "end date filter"
// @Success      200  {object}  response.AppSuccess
// @Failure      400  {object}  response.AppError
// @Failure      401  {object}  response.AppError
// @Failure      500  {object}  response.AppError
// @Failure      422  {object}  response.AppError
// @Security ApiKeyAuth
// @Router       /report-transaction-merchant [get]
func (instance *reportHandler) DoReportTransactionMerchant(c *fiber.Ctx) error {
	user := middleware.User(c.Context())
	if user == nil {
		return response.Response(c, response.New(fiber.StatusUnauthorized, response.WithMessage("Unauthorized")))
	}
	paginate, err := paginator.Paginate(c)
	if err != nil {
		return response.Response(c, response.New(fiber.StatusBadRequest, response.WithMessage(response.ErrBadRequest.Error())))

	}

	req := new(domain.ReportTransactionRequest)
	if err := c.QueryParser(req); err != nil {
		return response.Response(c, response.New(fiber.StatusBadRequest, response.WithMessage(response.ErrBadRequest.Error())))
	}

	req.Pagination = paginate

	report_transactions, pagination, err := instance.reportTransactionService.ReportTransactionMerchant(c.Context(), user.ID, *req)
	if err != nil {
		return response.Response(c, err)
	}

	return response.Success(c, fiber.StatusOK, response.SuccessData(report_transactions), response.SuccessMeta(pagination))
}

// Report Transaction Outlet godoc
// @Summary      Report Transaction Outlet
// @Description  API for get report transaction merchant
// @Tags         Report
// @Accept       json
// @Produce      json
// @Param        page   query     int  false  "page filter"  	   @Default("1")
// @Param        limit  query     int  false  "data limit filter"  @Default("10")
// @Param        start_date  query     string  false  "start date filter"
// @Param        end_date  query     string  false  "end date filter"
// @Success      200  {object}  response.AppSuccess
// @Failure      400  {object}  response.AppError
// @Failure      401  {object}  response.AppError
// @Failure      500  {object}  response.AppError
// @Failure      422  {object}  response.AppError
// @Security ApiKeyAuth
// @Router       /report-transaction-outlet/:outlet_id [get]
func (instance *reportHandler) DoReportTransactionOutlet(c *fiber.Ctx) error {
	user := middleware.User(c.Context())
	if user == nil {
		return response.Response(c, response.New(fiber.StatusUnauthorized, response.WithMessage("Unauthorized")))
	}
	paginate, err := paginator.Paginate(c)
	if err != nil {
		return response.Response(c, response.New(fiber.StatusBadRequest, response.WithMessage(response.ErrBadRequest.Error())))

	}

	outletID, err := strconv.ParseUint(c.Params("outlet_id"), 10, 64)
	if err != nil {
		return response.Response(c, response.New(fiber.StatusBadRequest, response.WithMessage(response.ErrBadRequest.Error())))
	}

	req := new(domain.ReportTransactionRequest)
	if err := c.QueryParser(req); err != nil {
		return response.Response(c, response.New(fiber.StatusBadRequest, response.WithMessage(response.ErrBadRequest.Error())))
	}

	req.Pagination = paginate
	req.OutletID = &outletID

	report_transactions, pagination, err := instance.reportTransactionService.ReportTransactionOutlet(c.Context(), user.ID, *req)
	if err != nil {
		return response.Response(c, err)
	}

	return response.Success(c, fiber.StatusOK, response.SuccessData(report_transactions), response.SuccessMeta(pagination))
}

