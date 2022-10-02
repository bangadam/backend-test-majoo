package auth_handler

import (
	"github.com/bangadam/backend-test-majoo/internal/domain"
	"github.com/bangadam/backend-test-majoo/internal/port"
	"github.com/bangadam/backend-test-majoo/pkg/response"
	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	app *fiber.App
	authService port.AuthService
}

func NewAuthHandler(app *fiber.App, authService port.AuthService)  {
	authHandler := &authHandler{
		app: app,
		authService: authService,
	}
	
	api := authHandler.app.Group("/api/v1")
	api.Post("/login", authHandler.DoLogin)
}

// Do Login godoc
// @Summary      Login
// @Description  API for login and get token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        user  body      domain.LoginRequest  true  "Login Account"
// @Success      200  {object}  response.AppSuccess
// @Failure      400  {object}  response.AppError
// @Failure      401  {object}  response.AppError
// @Failure      500  {object}  response.AppError
// @Failure      422  {object}  response.AppError
// @Router       /login [post]
func (instance *authHandler) DoLogin(c *fiber.Ctx) error {
	req := new(domain.LoginRequest)
	if err := c.BodyParser(req); err != nil {
		return response.Response(c, response.New(fiber.StatusBadRequest, response.WithMessage(response.ErrBadRequest.Error())))
	}

	if err := req.LoginValidation(); err != nil {
		return response.Response(c, response.New(fiber.StatusUnprocessableEntity, response.WithMeta(err)))
	}

	result, err := instance.authService.Login(c.Context(), *req)
	if err != nil {
		return response.Response(c, err)
	}

	return response.Success(c, fiber.StatusOK, response.SuccessMeta(result))
}