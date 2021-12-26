package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func getUserHandler(c echo.Context) (err error) {

	return c.JSON(http.StatusOK, map[string]string{"sucesso": "Cadastro realizado com sucesso"})
}

func postUserHandler(c echo.Context) (err error) {

	user := new(User)

	if err = c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}
