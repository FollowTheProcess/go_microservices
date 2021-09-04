package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	logger *log.Logger
}

func NewHello(logger *log.Logger) *Hello {
	return &Hello{logger: logger}
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Println("Hello World")
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Oops", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	fmt.Fprintf(w, "Data: %s\n", data)
}
