package handle

import (
	"net/http"

	"regulus/app/gateway/repositories"
	"regulus/app/usecases/maintenance/master/maker"

	"github.com/labstack/echo"
)

/*

MakerComplexSearchItems メーカー検索時に必要な検索項目を取得してJSONで送信する

*/
func MakerComplexSearchItems(c echo.Context) error {
	groupRepo := repositories.NewGroupRepo()
	result, err := maker.ComplexSearchItems(groupRepo)
	if err == nil {
		return c.JSON(http.StatusOK, result)
	}
	return err
}
