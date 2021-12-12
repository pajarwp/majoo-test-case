package entity

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"majoo-test-case/config"
	"majoo-test-case/entity/user/repository"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func AuthenticationMiddleware(r repository.UserRepository) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := strings.Split(c.Request().Header.Get("Authorization"), " ")
			if len(token) < 2 {
				return c.JSON(http.StatusUnauthorized, "Authorization Token Required")
			}

			if token[0] != "Bearer" {
				return c.JSON(http.StatusUnauthorized, "Authorization Token Required")
			}

			claims := jwt.MapClaims{}
			_, err := jwt.ParseWithClaims(token[1], claims, func(*jwt.Token) (interface{}, error) {
				return []byte(config.JWT_SECRET), nil
			})
			if err != nil {
				return c.JSON(http.StatusUnauthorized, "Authorization Token Invalid")
			}
			userIdClaims := claims["id"].(float64)

			json_map := make(map[string]interface{})
			err = json.NewDecoder(c.Request().Body).Decode(&json_map)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err)
			}

			if val, ok := json_map["merchant_id"]; ok {
				merchantId := val.(float64)
				userId, err := r.GetUserIDByMerchant(int(merchantId))
				if err != nil {
					return c.JSON(http.StatusInternalServerError, err)
				}
				if userId != int(userIdClaims) {
					return c.JSON(http.StatusForbidden, "Invalid  Access to Merchant")
				}
			}

			if val, ok := json_map["outlet_id"]; ok {
				outletId := val.(float64)
				userId, err := r.GetUserIDByOutlet(int(outletId))
				if err != nil {
					return c.JSON(http.StatusInternalServerError, err)
				}
				if userId != int(userIdClaims) {
					return c.JSON(http.StatusForbidden, "Invalid  Access to Outlet")
				}
			}

			body, err := json.Marshal(json_map)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err)
			}
			c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(body))

			return next(c)
		}
	}
}
