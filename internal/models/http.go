package models

type ListProduct struct {
	ID            string
	Name          string
	Description   string
	ThumbnailURL  string
	OverallRating float32
}

type Products struct {
	Products []ListProduct
	NextPage *int
}
