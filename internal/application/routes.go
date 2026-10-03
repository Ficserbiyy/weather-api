package application

import (
	"net/http"

	"github.com/Ficserbiyy/weather-api/internal/handler"
	"github.com/Ficserbiyy/weather-api/internal/repository/weather"
	"github.com/gin-gonic/gin"
)

func (a *App) loadRoutes() {
	// Gin router with default middleware
	router := gin.Default()

	router.HEAD("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	a.loadWeatherRoutes(router.Group("/weather"))

	a.router = router
}

func (a *App) loadWeatherRoutes(router *gin.RouterGroup) {
	weatherHandler := &handler.WeatherHandler{
		Repo: &weather.RedisRepo{
			Client: a.rdb,
		},
		APIKey: a.cfg.APIKey,
	}

	router.GET("/", weatherHandler.Homepage())
}
