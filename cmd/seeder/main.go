package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/schollz/progressbar/v3"
	mongodb "github.com/shackra/reviewer/internal/repository/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	PRODUCTS = 100_000
)

func main() {
	faker := gofakeit.New(0)

	connectionString := os.Getenv("MONGODB_URL")

	clientOptions := options.Client().ApplyURI(connectionString)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	repo := mongodb.New(client)

	bar := progressbar.Default(PRODUCTS)
	for i := 0; i < PRODUCTS; i++ {
		err := repo.AddProduct(
			context.TODO(),
			faker.ProductName(),
			faker.ProductDescription(),
			"",
		)
		if err != nil {
			log.Fatal(err)
		}
		bar.Add(1)
	}
}
