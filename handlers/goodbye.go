package handlers

import (
	"log"
	"net/http"
)

type GoodBye struct {
	logger *log.Logger
}

func NewGoodBye(logger *log.Logger) *GoodBye {
	return &GoodBye{logger: logger}
}

func (g *GoodBye) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g.logger.Println("Goodbye World")
	w.Write([]byte("Goodbye"))
}
