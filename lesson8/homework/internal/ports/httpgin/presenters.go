package httpgin

import (
	"homework8/internal/ads"
)

type Response struct {
	Data  adResponse `json:"data"`
	Error string     `json:"error"`
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

type changeAdStatusRequest struct {
	Published bool  `json:"published"`
	UserID    int64 `json:"user_id"`
}

type updateAdRequest struct {
	Title  string `json:"title"`
	Text   string `json:"text"`
	UserID int64  `json:"user_id"`
}

func AdSuccessResponse(ad *ads.Ad) Response {
	return Response{
		Data: adResponse{
			ID:        ad.ID,
			Title:     ad.Title,
			Text:      ad.Text,
			AuthorID:  ad.AuthorID,
			Published: ad.Published,
		},
	}
}

func AdErrorResponse(err error) Response {
	return Response{
		Error: err.Error(),
	}
}
