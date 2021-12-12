package delivery

import (
	"majoo-test-case/entity"
	"majoo-test-case/entity/user"
	"majoo-test-case/entity/user/usecase"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type UserHttpDelivery struct {
	UserUsecase usecase.UserUsecase
}

func NewHttpDelivery(e *echo.Echo, u usecase.UserUsecase) {
	handler := &UserHttpDelivery{
		UserUsecase: u,
	}
	e.POST("/user/login", handler.UserLogin)
}

func (u *UserHttpDelivery) UserLogin(c echo.Context) error {
	payload := new(user.UserLoginModel)
	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, entity.BadRequestResponse())
	}
	if err := validator.New().Struct(payload); err != nil {
		return c.JSON(http.StatusBadRequest, entity.BadRequestResponse())
	}

	token, err := u.UserUsecase.UserLogin(payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, entity.InternalServerErrorResponse(err.Error()))
	}
	data := map[string]interface{}{
		"token": token,
	}
	return c.JSON(http.StatusOK, entity.NewSuccessResponse(data))
}
