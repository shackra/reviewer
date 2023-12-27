package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type service interface {
	ListProducts() (interface{}, error) // TBD
}

var ErrWriteResponse = errors.New("problem was encountered when responding to the user")

type Server struct {
	productService service
}

func respond(w http.ResponseWriter, message interface{}) error {
	response, err := json.Marshal(message)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(response)
	return errors.Join(ErrWriteResponse, err)
}

func (s *Server) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := s.productService.ListProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = respond(w, products)
	if err != nil && !errors.Is(err, ErrWriteResponse) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		log.Printf("GetProducts: %v", err)
	}
}
