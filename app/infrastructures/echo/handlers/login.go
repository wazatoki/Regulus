package handlers

import (
	"net/http"

	"regulus/app/repositories"
	"regulus/app/usecases"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

/*
Login return staff and jwt token
*/
func Login(c echo.Context) error {
	m := echo.Map{}
	c.Bind(&m)

	repo := repositories.NewStaffRepo()

	staff, token, err := usecases.Login(repo, m["id"].(string), m["password"].(string))

	if err != nil {
		return c.JSON(http.StatusUnauthorized, "")
	}

	staff.Password = ""

	return c.JSON(http.StatusOK, echo.Map{"staff": staff, "jwtToken": token})
}

func getAuthStaffID(c echo.Context) string {
	if c.Get("user") == nil {
		return ""
	}
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	staffID := claims["staffID"].(string)
	return staffID
}
