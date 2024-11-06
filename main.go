package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"path/filepath"
)

func main() {
	server := echo.New()

	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, "X-App-Token", "Authorization"},
	}))

	server.GET("/:folder/:filename", func(c echo.Context) error {
		folder := c.Param("folder")
		filename := c.Param("filename")

		filePath := filepath.Join(folder, filename)
		c.Response().Header().Set(echo.HeaderContentType, "audio/wav")

		if err := c.File(filePath); err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "File not found"})
		}
		return nil
	})

	if err := server.Start(fmt.Sprintf(":%d", 7070)); err != nil {
		server.Logger.Info("shutting down the server")
	}
}
