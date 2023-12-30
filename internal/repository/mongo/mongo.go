package mongodb

import (
	"context"

	"github.com/shackra/reviewer/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	database       = "reviewdb"
	collectionName = "product"
)

type Mongo struct {
	client *mongo.Client
}

func New(client *mongo.Client) *Mongo {
	return &Mongo{
		client: client,
	}
}

func (m *Mongo) GetProducts(page, size int) ([]models.Product, bool, error) {
	var products []models.Product

	collection := m.client.Database(database).Collection(collectionName)

	total, err := collection.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		return nil, false, err
	}

	totalPages := int(total) / size
	// handle rounding
	if int(total)%size != 0 {
		totalPages++
	}

	findOptions := options.Find()
	findOptions.SetSkip(int64((page - 1) * size))
	findOptions.SetLimit(int64(size))

	cursor, err := collection.Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		return nil, false, err
	}

	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var product models.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, false, err
		}
		products = append(products, product)
	}

	return products, totalPages-int(total) > 0, nil
}

func (m *Mongo) AddProductReview(id, reviewer, text string, rating float32) error {
	collection := m.client.Database(database).Collection(collectionName)

	_, err := collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": id},
		bson.M{
			"$push": bson.M{
				"reviews": models.Review{
					Name:   reviewer,
					Text:   text,
					Rating: rating,
				},
			},
		},
	)
	if err != nil {
		return err
	}

	return nil
}
