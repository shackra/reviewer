package http

type Product struct {
	ID            string
	Name          string
	Description   string
	ThumbnailURL  string
	OverallRating float32
}

type ListProducts struct {
	Products []Product
	NextPage *string
}
