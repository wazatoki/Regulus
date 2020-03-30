package handle

import (
	"net/http"

	"regulus/app/usecases/maintenance/master/maker"

	"github.com/labstack/echo"
)

/*

MakerComplexSearchItems メーカー検索時に必要な検索項目を取得してJSONで送信する

*/
func MakerComplexSearchItems(c echo.Context) error {
	return c.JSON(http.StatusOK, maker.ComplexSearchItems)
}
