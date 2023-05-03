package httpgin

import (
	"homework10/internal/ads"
)

type response struct {
	Data  any    `json:"data"`
	Error string `json:"error"`
}

type userResponse struct {
	ID       int64  `json:"id"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
}

type createUserRequest struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
}

type updateUserRequest struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
}

type deleteUserRequest struct {
	UserID int64 `json:"user_id"`
}

type createAdRequest struct {
	Title  string `json:"title"`
	Text   string `json:"text"`
	UserID int64  `json:"user_id"`
}

type adResponse struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Text      string `json:"text"`
	AuthorID  int64  `json:"author_id"`
	Published bool   `json:"published"`
}

type adsResponse struct {
	Data []adResponse `json:"data"`
}

type changeAdStatusRequest struct {
	Published bool  `json:"published"`
	UserID    int64 `json:"user_id"`
}

type deleteAdRequest struct {
	AdID   int64 `json:"ad_id"`
	UserID int64 `json:"user_id"`
}

type updateAdRequest struct {
	Title  string `json:"title"`
	Text   string `json:"text"`
	UserID int64  `json:"user_id"`
}

func userSuccessResponse(user *ads.User) response {
	return response{
		Data: userResponse{
			ID:       user.ID,
			Nickname: user.Nickname,
			Email:    user.Email,
		},
	}
}

func adSuccessResponse(ad *ads.Ad) response {
	return response{
		Data: adToAdResponse(*ad),
	}
}

func adsSuccessResponse(ads []*ads.Ad) adsResponse {
	result := adsResponse{
		Data: make([]adResponse, len(ads)),
	}
	for i, ad := range ads {
		result.Data[i] = adToAdResponse(*ad)
	}
	return result
}

func errorResponse(err error) response {
	return response{
		Error: err.Error(),
	}
}

func adToAdResponse(ad ads.Ad) adResponse {
	return adResponse{
		ID:        ad.ID,
		Title:     ad.Title,
		Text:      ad.Text,
		AuthorID:  ad.AuthorID,
		Published: ad.Published,
	}
}
