package handlers

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Gprisco/decanto-golang/services"
)

type Hello struct {
	logger *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) HandleGet(rw http.ResponseWriter, r *http.Request) {
	h.logger.Println("Hit Hello World handler")

	d, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	rw.Write([]byte(services.HelloWorld(d)))
}
