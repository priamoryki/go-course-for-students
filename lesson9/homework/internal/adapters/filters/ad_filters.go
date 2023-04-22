package filters

import "homework9/internal/ads"

func NewFilterNonPublished() Filter[ads.Ad] {
	return DefaultFilter[ads.Ad]{
		condition: func(ad *ads.Ad) bool {
			return ad.Published
		},
	}
}

func NewFilterByAuthor() Filter[ads.Ad] {
	return SortFilter[ads.Ad]{
		comparator: func(ad1 *ads.Ad, ad2 *ads.Ad) bool {
			return ad1.AuthorID < ad2.AuthorID
		},
	}
}

func NewFilterByCreationTime() Filter[ads.Ad] {
	return SortFilter[ads.Ad]{
		comparator: func(ad1 *ads.Ad, ad2 *ads.Ad) bool {
			return ad2.CreationTime.After(ad1.CreationTime)
		},
	}
}
