package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	now := time.Now()

	volumeID := os.Getenv("RLWY_INTERNAL_VOLUME_ID")
	mountPath := os.Getenv("RLWY_INTERNAL_VOLUME_MOUNT_PATH")

	if volumeID != "" && mountPath != "" {
		_ = os.WriteFile(filepath.Join(mountPath, fmt.Sprint(now.Unix())), []byte(now.UTC().String()), 0644)
	}
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
	e.Static("/data", "/railway/data")
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   mountPath,
		Browse: true,
	}))
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
