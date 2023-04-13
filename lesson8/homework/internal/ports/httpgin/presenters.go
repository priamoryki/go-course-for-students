package httpgin

import (
	"homework8/internal/ads"
)

type response struct {
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

type adsResponse struct {
	Data []adResponse `json:"data"`
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

func adErrorResponse(err error) response {
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
