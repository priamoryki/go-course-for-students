package httpgin

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	"homework8/internal/app"
)

// Метод для создания пользователя (user)
func createUser(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqBody createUserRequest
		err := c.BindJSON(&reqBody)
		if err != nil {
			c.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		user, err := a.CreateUser(reqBody.Nickname, reqBody.Email)
		if err != nil {
			c.JSON(getStatusByError(err), errorResponse(err))
			return
		}

		c.JSON(http.StatusOK, userSuccessResponse(user))
	}
}

// Метод для обновления пользователя (user)
func updateUser(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqBody updateUserRequest
		if err := c.BindJSON(&reqBody); err != nil {
			c.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		user, err := a.UpdateUser(userID, reqBody.Nickname, reqBody.Email)
		if err != nil {
			c.JSON(getStatusByError(err), errorResponse(err))
			return
		}

		c.JSON(http.StatusOK, userSuccessResponse(user))
	}
}

// Метод для поиска пользователя (user)
func findUser(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		searchQuery := c.Query("search_query")
		if searchQuery == "" {
			c.JSON(http.StatusBadRequest, errorResponse(errors.New("search_query parameter is empty")))
			return
		}

		user, err := a.FindUser(searchQuery)
		if err != nil {
			c.JSON(getStatusByError(err), errorResponse(err))
			return
		}

		c.JSON(http.StatusOK, userSuccessResponse(user))
	}
}

// Метод получения объявлений (ads)
func listAds(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		// С фронта приходит битовая маска фильтров
		filters := c.Query("filters")
		bitmask, err := strconv.ParseInt(filters, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		c.JSON(http.StatusOK, adsSuccessResponse(a.ListAds(bitmask)))
	}
}

// Метод для создания объявления (ad)
func createAd(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqBody createAdRequest
		err := c.BindJSON(&reqBody)
		if err != nil {
			c.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		ad, err := a.CreateAd(reqBody.Title, reqBody.Text, reqBody.UserID)
		if err != nil {
			c.JSON(getStatusByError(err), errorResponse(err))
			return
		}

		c.JSON(http.StatusOK, adSuccessResponse(ad))
	}
}

// Метод для изменения статуса объявления (опубликовано - Published = true или снято с публикации Published = false)
func changeAdStatus(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqBody changeAdStatusRequest
		if err := c.BindJSON(&reqBody); err != nil {
			c.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		adID, err := strconv.ParseInt(c.Param("ad_id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		ad, err := a.ChangeAdStatus(adID, reqBody.UserID, reqBody.Published)
		if err != nil {
			c.JSON(getStatusByError(err), errorResponse(err))
			return
		}

		c.JSON(http.StatusOK, adSuccessResponse(ad))
	}
}

// Метод для получения объявления (ad)
func getAd(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		adID, err := strconv.ParseInt(c.Param("ad_id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		ad, err := a.GetAd(adID)
		if err != nil {
			c.JSON(getStatusByError(err), errorResponse(err))
			return
		}

		c.JSON(http.StatusOK, adSuccessResponse(ad))
	}
}

// Метод для обновления текста(Text) или заголовка(Title) объявления
func updateAd(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqBody updateAdRequest
		if err := c.BindJSON(&reqBody); err != nil {
			c.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		adID, err := strconv.ParseInt(c.Param("ad_id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		ad, err := a.UpdateAd(adID, reqBody.UserID, reqBody.Title, reqBody.Text)
		if err != nil {
			c.JSON(getStatusByError(err), errorResponse(err))
			return
		}

		c.JSON(http.StatusOK, adSuccessResponse(ad))
	}
}

// Метод для поиска объявления (ad)
func findAd(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		searchQuery := c.Query("search_query")
		if searchQuery == "" {
			c.JSON(http.StatusBadRequest, errorResponse(errors.New("search_query parameter is empty")))
			return
		}

		ad, err := a.FindAd(searchQuery)
		if err != nil {
			c.JSON(getStatusByError(err), errorResponse(err))
			return
		}

		c.JSON(http.StatusOK, adSuccessResponse(ad))
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
