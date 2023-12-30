package products

import (
	"errors"
	"testing"

	"github.com/shackra/reviewer/internal/models"
	"github.com/subtle-byte/mockigo/match"
)

func TestListProducts(t *testing.T) {
	repo := NewRepositoryMock(t)

	repo.EXPECT().GetProducts(match.Eq(1), match.Eq(10)).Return([]models.Product{
		{
			ID:          "",
			Name:        "",
			Description: "",
			ImgURL:      "",
			Reviews: []models.Review{
				{
					ID:     "",
					Name:   "",
					Text:   "",
					Rating: 3,
				},
				{
					ID:     "",
					Name:   "",
					Text:   "",
					Rating: 3,
				},
				{
					ID:     "",
					Name:   "",
					Text:   "",
					Rating: 5,
				},
			},
		},
		{
			ID:          "",
			Name:        "",
			Description: "",
			ImgURL:      "",
			Reviews: []models.Review{
				{
					ID:     "",
					Name:   "",
					Text:   "",
					Rating: 4,
				},
				{
					ID:     "",
					Name:   "",
					Text:   "",
					Rating: 4,
				},
				{
					ID:     "",
					Name:   "",
					Text:   "",
					Rating: 1,
				},
			},
		},
		{
			ID:          "",
			Name:        "",
			Description: "",
			ImgURL:      "",
			Reviews: []models.Review{
				{
					ID:     "",
					Name:   "",
					Text:   "",
					Rating: 5,
				},
				{
					ID:     "",
					Name:   "",
					Text:   "",
					Rating: 5,
				},
				{
					ID:     "",
					Name:   "",
					Text:   "",
					Rating: 1.4,
				},
			},
		},
	}, false, nil)

	service := &Service{
		mongo: repo,
	}

	result, err := service.ListProducts(1, 10)
	if err != nil {
		t.Errorf("got unexpected error %v", err)
	}

	if result == nil {
		t.Errorf("unexpected result: got %v want %v", result, &models.ListProduct{})
	}
}

func TestListProductsFail(t *testing.T) {
	repo := NewRepositoryMock(t)

	repo.EXPECT().
		GetProducts(match.Eq(1), match.Eq(10)).
		Return(nil, false, errors.New(`random error`))

	service := &Service{
		mongo: repo,
	}

	_, err := service.ListProducts(1, 10)
	if err == nil {
		t.Errorf("got unexpected error %v", err)
	}
}

func TestAddProductReview(t *testing.T) {
	repo := NewRepositoryMock(t)

	repo.EXPECT().
		AddProductReview(match.Eq("123"), match.Eq("Test User"), match.Eq("Lorem Ipsum"), match.Eq[float32](5)).
		Return(nil)

	service := &Service{
		mongo: repo,
	}

	err := service.AddReview("123", "Test User", "Lorem Ipsum", 5)
	if err != nil {
		t.Errorf("got unexpected error %v", err)
	}
}

func TestAddProductReviewFails(t *testing.T) {
	repo := NewRepositoryMock(t)

	repo.EXPECT().
		AddProductReview(match.Eq("123"), match.Eq("Test User"), match.Eq("Lorem Ipsum"), match.Eq[float32](5)).
		Return(errors.New(`random error`))

	service := &Service{
		mongo: repo,
	}

	err := service.AddReview("123", "Test User", "Lorem Ipsum", 5)
	if err == nil {
		t.Errorf("got unexpected error %v", err)
	}
}
