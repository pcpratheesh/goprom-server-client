package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		req, err := http.NewRequest("GET", os.Getenv("RECEIVER_HOST"), nil)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		req.Header.Set("X-From-Service", os.Getenv("APP_NAME"))

		// Make the request
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		defer resp.Body.Close()

		return c.String(http.StatusOK, "Request send from sender to receiver")
	})

	if err := e.Start(":8001"); err != nil {
		panic(err)
	}
}
