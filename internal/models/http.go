package models

type ListProduct struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	ThumbnailURL  string  `json:"thumbnail_url"`
	AverageRating float32 `json:"average_rating"`
}

type Products struct {
	Products []ListProduct `json:"products"`
	NextPage *int          `json:"next_page"`
}
