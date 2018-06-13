package main
import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "ok(v1)")
	})
	e.Logger.Fatal(e.Start(":8090"))
}