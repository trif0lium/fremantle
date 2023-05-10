package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]bool{"ok": true})
	})
	e.GET("/hostname", func(c echo.Context) error {
		hostname, err := os.Hostname()
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{"hostname": hostname})
	})
	e.GET("/git", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"RAILWAY_GIT_COMMIT_SHA": os.Getenv("RAILWAY_GIT_COMMIT_SHA"),
		})
	})
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
