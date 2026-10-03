package handler

import (
	"net/http"

	"github.com/Ficserbiyy/weather-api/internal/repository/weather"
	"github.com/gin-gonic/gin"
)

type WeatherHandler struct {
	Repo   *weather.RedisRepo
	APIKey string
}

func (h *WeatherHandler) Homepage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Status(http.StatusOK)
	}
}
