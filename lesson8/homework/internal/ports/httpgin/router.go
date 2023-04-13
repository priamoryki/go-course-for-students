package httpgin

import (
	"github.com/gin-gonic/gin"
	"homework8/internal/app"
)

func AppRouter(r gin.IRouter, a app.App) {
	r.PUT("/users", createUser(a))                 // Метод для создания пользователя (user)
	r.PUT("/users/:user_id", updateUser(a))        // Метод для обновления пользователя (user)
	r.GET("/users/find", findUser(a))              // Метод для поиска объявления (ad)
	r.GET("/ads", listAds(a))                      // Метод для получения объявлений (ads)
	r.POST("/ads", createAd(a))                    // Метод для создания объявления (ad)
	r.PUT("/ads/:ad_id/status", changeAdStatus(a)) // Метод для изменения статуса объявления (опубликовано - Published = true или снято с публикации Published = false)
	r.GET("/ads/:ad_id", getAd(a))                 // Метод для получения объявления (ad)
	r.PUT("/ads/:ad_id", updateAd(a))              // Метод для обновления текста(Text) или заголовка(Title) объявления
	r.GET("/ads/find", findAd(a))                  // Метод для поиска объявления (ad)
}
