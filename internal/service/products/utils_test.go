package products

import (
	"testing"

	"github.com/shackra/reviewer/internal/models"
)

func TestGetRating(t *testing.T) {
	input := []models.Review{
		{
			ID:     "",
			Name:   "",
			Text:   "",
			Rating: 2,
		},
		{
			ID:     "",
			Name:   "",
			Text:   "",
			Rating: 2,
		},
		{
			ID:     "",
			Name:   "",
			Text:   "",
			Rating: 5,
		},
	}

	rating := getRating(input)
	if rating != 3 {
		t.Errorf("unexpected result %f, wanted %f", rating, 3.0)
	}
}
