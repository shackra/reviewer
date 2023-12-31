package mongodb

import (
	"context"
	"fmt"

	"github.com/shackra/reviewer/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	database       = "app"
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

func (m *Mongo) GetProducts(ctx context.Context, page, size int) ([]models.Product, bool, error) {
	var products []models.Product

	collection := m.client.Database(database).Collection(collectionName)

	total, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, false, err
	}

	totalPages := int(total) / size
	// handle rounding
	if int(total)%size != 0 {
		totalPages++
	}

	pipeline := bson.A{
		bson.M{
			"$project": bson.M{
				"_id":            1,
				"name":           1,
				"description":    1,
				"img_url":        1,
				"average_rating": bson.M{"$avg": "$reviews.rating"},
			},
		},
	}

	pipeline = append(pipeline,
		bson.M{"$skip": (page - 1) * size},
		bson.M{"$limit": size},
	)

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, false, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var product models.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, false, err
		}
		products = append(products, product)
	}

	return products, totalPages-page > 0, nil
}

func (m *Mongo) AddProductReview(
	ctx context.Context,
	id, reviewer, text string,
	rating float64,
) error {
	collection := m.client.Database(database).Collection(collectionName)

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid ObjectID '%s'", id)
	}

	result, err := collection.UpdateOne(
		ctx,
		bson.D{{"_id", objID}},
		bson.D{
			{"$push", bson.D{
				{"reviews", models.Review{
					Name:   reviewer,
					Text:   text,
					Rating: rating,
				}},
			}},
		},
	)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("no document with _id '%s' found", id)
	}

	return nil
}

func (m *Mongo) AddProduct(ctx context.Context, name, description, image string) error {
	collection := m.client.Database(database).Collection(collectionName)

	newProduct := models.Product{
		Name:        name,
		Description: description,
		ImgURL:      image,
		Reviews:     []models.Review{},
	}

	result, err := collection.InsertOne(ctx, newProduct)
	if err != nil {
		return err
	}

	if result.InsertedID == "" {
		return fmt.Errorf("nothing was recorded")
	}

	return nil
}
