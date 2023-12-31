package http

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/shackra/reviewer/internal/models"
)

type Service interface {
	ListProducts(context.Context, int, int) (*models.Products, error)
	AddReview(context.Context, string, string, string, float32) error
}

type Server struct {
	productService Service
}

func New(service Service) *Server {
	return &Server{
		productService: service,
	}
}

type ErrorMessage struct {
	Status string `json:"status"`
	Reason string `json:"reason"`
}

func (e *ErrorMessage) Message() string {
	message, _ := json.Marshal(e)

	return string(message)
}

func (e *ErrorMessage) Error() string {
	return e.Message()
}

func (e *ErrorMessage) GoString() string {
	return fmt.Sprintf(`Status: "%s", Reason: "%v"`, e.Status, e.Reason)
}

func newErrorMessage(err error) *ErrorMessage {
	return &ErrorMessage{
		Status: "error",
		Reason: err.Error(),
	}
}

func newOkMessage() *ErrorMessage {
	return &ErrorMessage{Status: "ok"}
}

func pagination(page, amount string) (int, int, error) {
	if page == "" {
		page = "1"
	}

	if amount == "" {
		amount = "10"
	}

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return 0, 0, err
	}

	amountInt, err := strconv.Atoi(amount)
	if err != nil {
		return 0, 0, err
	}

	return pageInt, amountInt, nil
}

func (s *Server) GetProducts(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	page, amount, err := pagination(query.Get("page"), query.Get("amount"))
	if err != nil {
		err = newErrorMessage(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	products, err := s.productService.ListProducts(ctx, page, amount)
	if err != nil {
		err = newErrorMessage(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := json.Marshal(products)
	if err != nil {
		err = newErrorMessage(fmt.Errorf("cannot serialize data: %v", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(response)
	if err != nil {
		log.Printf("GetProducts: %v", err)
	}
}

func (s *Server) AddReview(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID := vars["id"]

	var p AddProductReviewRequest
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		err = newErrorMessage(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = validate(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	err = s.productService.AddReview(ctx, productID, p.Name, p.Text, p.Rating)
	if err != nil {
		statusError := http.StatusInternalServerError
		if strings.Contains(err.Error(), "no document with _id") {
			statusError = http.StatusNotFound
		}
		if strings.Contains(err.Error(), "invalid ObjectID") {
			statusError = http.StatusBadRequest
		}
		err = newErrorMessage(err)
		http.Error(w, err.Error(), statusError)
		return
	}

	fmt.Fprint(w, newOkMessage().Message())
}
