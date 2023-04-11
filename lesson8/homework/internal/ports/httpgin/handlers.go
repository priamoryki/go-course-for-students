package httpgin

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	"homework8/internal/app"
)

// Метод для создания объявления (ad)
func createAd(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqBody createAdRequest
		err := c.BindJSON(&reqBody)
		if err != nil {
			c.JSON(http.StatusBadRequest, AdErrorResponse(err))
			return
		}

		ad, err := a.CreateAd(reqBody.Title, reqBody.Text, reqBody.UserID)
		if err != nil {
			c.JSON(getStatusByError(err), AdErrorResponse(err))
			return
		}

		c.JSON(http.StatusOK, AdSuccessResponse(ad))
	}
}

// Метод для изменения статуса объявления (опубликовано - Published = true или снято с публикации Published = false)
func changeAdStatus(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqBody changeAdStatusRequest
		if err := c.BindJSON(&reqBody); err != nil {
			c.JSON(http.StatusBadRequest, AdErrorResponse(err))
			return
		}

		adID, err := strconv.Atoi(c.Param("ad_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, AdErrorResponse(err))
			return
		}

		ad, err := a.ChangeAdStatus(int64(adID), reqBody.UserID, reqBody.Published)
		if err != nil {
			c.JSON(getStatusByError(err), AdErrorResponse(err))
			return
		}

		c.JSON(http.StatusOK, AdSuccessResponse(ad))
	}
}

// Метод для обновления текста(Text) или заголовка(Title) объявления
func updateAd(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqBody updateAdRequest
		if err := c.BindJSON(&reqBody); err != nil {
			c.JSON(http.StatusBadRequest, AdErrorResponse(err))
			return
		}

		adID, err := strconv.Atoi(c.Param("ad_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, AdErrorResponse(err))
			return
		}

		ad, err := a.UpdateAd(int64(adID), reqBody.UserID, reqBody.Title, reqBody.Text)
		if err != nil {
			c.JSON(getStatusByError(err), AdErrorResponse(err))
			return
		}

		c.JSON(http.StatusOK, AdSuccessResponse(ad))
	}
}

func getStatusByError(err error) int {
	switch {
	case errors.Is(err, app.ErrNotUsersAd):
		return http.StatusForbidden
	case errors.Is(err, app.ErrValidation):
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
