package products

import "github.com/shackra/reviewer/internal/models"

func getRating(reviews []models.Review) float32 {
	if len(reviews) == 0 {
		return 0
	}

	var overall float32 = 0
	for _, value := range reviews {
		overall += value.Rating
	}
	return overall / float32(len(reviews))
}
