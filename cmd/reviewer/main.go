package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	mongodb "github.com/shackra/reviewer/internal/repository/mongo"
	"github.com/shackra/reviewer/internal/service/products"
	transport "github.com/shackra/reviewer/internal/transport/http"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Define la cadena de conexión de MongoDB
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
	service := products.New(repo)
	server := transport.New(service)

	rr := mux.NewRouter()

	rr.HandleFunc("/product/{id}/review", server.AddReview).Methods(http.MethodPost)
	rr.HandleFunc("/product", server.GetProducts).Methods(http.MethodGet)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	err = http.ListenAndServe(":"+port, rr)

	if err != nil {
		log.Printf("server exited with error %v", err)
	}
}
