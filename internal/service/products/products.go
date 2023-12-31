package products

import (
	"context"

	"github.com/shackra/reviewer/internal/models"
)

type Service struct {
	mongo Repository
}

type Repository interface {
	GetProducts(ctx context.Context, page, size int) ([]models.Product, bool, error)
	AddProductReview(ctx context.Context, id, reviewer, text string, rating float64) error
}

func New(db Repository) *Service {
	return &Service{
		mongo: db,
	}
}

func (s *Service) ListProducts(ctx context.Context, page, amount int) (*models.Products, error) {
	pds, nextPage, err := s.mongo.GetProducts(ctx, page, amount)
	if err != nil {
		return nil, err
	}

	list := &models.Products{}
	if nextPage {
		nextPage := page + 1
		list.NextPage = &nextPage
	}

	for _, product := range pds {
		list.Products = append(list.Products, models.ListProduct{
			ID:            product.ID,
			Name:          product.Name,
			Description:   product.Description,
			ThumbnailURL:  product.ImgURL,
			AverageRating: product.AverageRating,
		})
	}

	return list, nil
}

func (s *Service) AddReview(ctx context.Context, id, reviewer, text string, rating float64) error {
	err := s.mongo.AddProductReview(ctx, id, reviewer, text, rating)
	if err != nil {
		return err
	}

	return nil
}
