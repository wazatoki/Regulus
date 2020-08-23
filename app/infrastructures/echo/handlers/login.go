package handlers

import (
	"net/http"

	"regulus/app/repositories"
	"regulus/app/usecases/login"

	"github.com/labstack/echo"
)

/*
Login return staff and jwt token
*/
func Login(c echo.Context) error {
	m := echo.Map{}
	c.Bind(&m)

	repo := repositories.NewStaffRepo()

	staff, token, err := login.Login(repo, m["id"].(string), m["password"].(string))

	if err != nil {
		return c.JSON(http.StatusUnauthorized, "")
	}

	return c.JSON(http.StatusOK, echo.Map{"staff": staff, "jwtToken": token})
}
