package http

type AddProductReviewRequest struct {
	Name   string  `validate:"required"   json:"name"`
	Text   string  `                      json:"text"`
	Rating float32 `validate:"gt=0,lte=5" json:"rating"`
}
