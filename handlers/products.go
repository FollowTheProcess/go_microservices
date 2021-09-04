package handlers

import (
	"log"
	"net/http"

	"github.com/FollowTheProcess/go_microservices/data"
)

type Products struct {
	logger *log.Logger
}

func NewProducts(logger *log.Logger) *Products {
	return &Products{logger: logger}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(w, r)
		return
	}

	// Catch non-allowed methods
	w.WriteHeader(http.StatusMethodNotAllowed)
}

// getProducts handles a HTTP GET to our products handler.
func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	products := data.GetProducts()
	err := products.ToJSON(w)
	if err != nil {
		http.Error(w, "could not serialize products to JSON", http.StatusInternalServerError)
		return
	}
}
