package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type api struct {
	router http.Handler
}

type Server interface {
	Router() http.Handler
}

func New() Server {
	a := &api{}

	r := mux.NewRouter()
	r.HandleFunc("/holamundo", a.holaMundo).Methods(http.MethodGet)
	a.router = r
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}

func (a *api) holaMundo(w http.ResponseWriter, r *http.Request) {
	log.Print("[INFO] init: HelloWorld()")
	defer func() {
		err := r.Body.Close()
		if err != nil {
			Response(struct {
				Message string `json:"message"`
			}{err.Error()}, http.StatusInternalServerError, w)
		}
	}()

	Response(struct {
		Message string `json:"message"`
	}{fmt.Sprint("Hola Mundo")}, http.StatusOK, w)

}

func Response(resp interface{}, statusCode int, w http.ResponseWriter) {
	response, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic("error")
	}
	w.WriteHeader(statusCode)
	_, _ = w.Write(response)
}
