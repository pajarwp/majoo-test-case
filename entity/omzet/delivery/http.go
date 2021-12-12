package delivery

import (
	"majoo-test-case/entity"
	"majoo-test-case/entity/omzet"
	"majoo-test-case/entity/omzet/usecase"
	"majoo-test-case/entity/user/repository"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type OmzetHttpDelivery struct {
	OmzetUsecase   usecase.OmzetUsecase
	UserRepository repository.UserRepository
}

func NewHttpDelivery(e *echo.Echo, u usecase.OmzetUsecase, r repository.UserRepository) {
	handler := &OmzetHttpDelivery{
		OmzetUsecase:   u,
		UserRepository: r,
	}
	omzet := e.Group("omzet")
	omzet.Use(entity.AuthenticationMiddleware(r))
	omzet.POST("/merchant", handler.GetMerchantOmzet)
	omzet.POST("/outlet", handler.GetOutletOmzet)
}

func (o *OmzetHttpDelivery) GetMerchantOmzet(c echo.Context) error {
	payload := new(omzet.MerchantOmzetRequest)
	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, entity.BadRequestResponse())
	}
	if err := validator.New().Struct(payload); err != nil {
		return c.JSON(http.StatusBadRequest, entity.BadRequestResponse())
	}

	resp, err := o.OmzetUsecase.GetMerchantOmzet(payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, entity.InternalServerErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, entity.NewSuccessResponse(resp))
}

func (o *OmzetHttpDelivery) GetOutletOmzet(c echo.Context) error {
	payload := new(omzet.OutletOmzetRequest)
	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, entity.BadRequestResponse())
	}
	if err := validator.New().Struct(payload); err != nil {
		return c.JSON(http.StatusBadRequest, entity.BadRequestResponse())
	}

	resp, err := o.OmzetUsecase.GetOutletOmzet(payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, entity.InternalServerErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, entity.NewSuccessResponse(resp))
}
