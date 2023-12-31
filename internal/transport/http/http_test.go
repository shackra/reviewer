package http

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/shackra/reviewer/internal/models"
	"github.com/subtle-byte/mockigo/match"
)

func TestListProducts(t *testing.T) {
	req, err := http.NewRequest("GET", "/product", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	service := NewServiceMock(t)
	service.EXPECT().
		ListProducts(match.AnyCtx(), match.Arg[int](match.Eq[int](1)), match.Arg[int](match.Eq[int](10))).
		Return(&models.Products{}, nil)

	api := &Server{
		productService: service,
	}

	handler := http.HandlerFunc(api.GetProducts)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"products":null,"next_page":null}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestListProductsFails(t *testing.T) {
	req, err := http.NewRequest("GET", "/product", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	service := NewServiceMock(t)
	service.EXPECT().
		ListProducts(match.AnyCtx(), match.Arg[int](match.Eq[int](1)), match.Arg[int](match.Eq[int](10))).
		Return(nil, errors.New(`random error`))

	api := &Server{
		productService: service,
	}

	handler := http.HandlerFunc(api.GetProducts)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf(
			"handler returned wrong status code: got %v want %v",
			status,
			http.StatusBadRequest,
		)
	}

	expected := `{"status":"error","reason":"random error"}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf(
			"handler returned unexpected body: got '%v' want '%v'",
			strings.TrimSpace(rr.Body.String()),
			expected,
		)
	}
}

func TestAddReview(t *testing.T) {
	service := NewServiceMock(t)
	service.EXPECT().
		AddReview(match.AnyCtx(), match.Eq("123"), match.Eq("Test Name"), match.Eq("Test review text"), match.Eq[float64](5)).
		Return(nil)

	api := &Server{
		productService: service,
	}

	r := mux.NewRouter()
	r.HandleFunc("/product/{id}/review", api.AddReview)

	// Crea una solicitud HTTP de prueba con el cuerpo JSON
	body := `{"name": "Test Name", "text": "Test review text", "rating": 5}`
	req, err := http.NewRequest("POST", "/product/123/review", strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{"id": "123"})

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"status":"ok","reason":""}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf(
			"handler returned unexpected body: got '%v' want '%v'",
			strings.TrimSpace(rr.Body.String()),
			expected,
		)
	}
}

func TestAddReviewFails(t *testing.T) {
	service := NewServiceMock(t)
	service.EXPECT().
		AddReview(match.AnyCtx(), match.Eq("123"), match.Eq("Test Name"), match.Eq("Test review text"), match.Eq[float64](5)).
		Return(errors.New(`random error`))

	api := &Server{
		productService: service,
	}

	r := mux.NewRouter()
	r.HandleFunc("/product/{id}/review", api.AddReview)

	// Crea una solicitud HTTP de prueba con el cuerpo JSON
	body := `{"name": "Test Name", "text": "Test review text", "rating": 5}`
	req, err := http.NewRequest("POST", "/product/123/review", strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{"id": "123"})

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"status":"error","reason":"random error"}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf(
			"handler returned unexpected body: got '%v' want '%v'",
			strings.TrimSpace(rr.Body.String()),
			expected,
		)
	}
}
