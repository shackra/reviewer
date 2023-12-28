package products

import "github.com/shackra/reviewer/internal/models"

type Service struct {
	mongo Repository
}

type Repository interface {
	GetProducts(page, size int) ([]models.Product, bool, error)
	AddProductReview(id, reviewer, text string, rating float32) error
}

func (s *Service) ListProducts(page, amount int) (*models.Products, error) {
	pds, nextPage, err := s.mongo.GetProducts(page, amount)
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
			OverallRating: getRating(product.Reviews),
		})
	}

	return list, nil
}

func (s *Service) AddProductReview(id, reviewer, text string, rating float32) error {
	err := s.mongo.AddProductReview(id, reviewer, text, rating)
	if err != nil {
		return err
	}

	return nil
}
