package core

import (
	"encoding/json"
	"fmt"
	"gys/pkg"
	"log"
	"net/http"
	"sync"
)

type Handler struct {
	sync.Mutex
}


//API
func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	//process response
	//another possibility var v interface {} but after that printing password does not work
	var v pkg.GysServer
	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		fmt.Println(err)
	}
	log.Println(v)
	//response := Response{Email: "blab@blab", Age: 5}
	//jsonBytes, err := json.Marshal(response)
	//if err != nil {
	//	w.WriteHeader(http.StatusInternalServerError)
	//	w.Write([]byte(err.Error()))
	//}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"success":true}`))
}

func newHandler() *Handler{
	return &Handler{
	}

}

func Server() {
	responseHandlers := newHandler()
	//http.HandleFunc("/coasters", coasterHandlers.coasters)
	http.HandleFunc("/api/v1/scrap/", responseHandlers.get)
	//http.HandleFunc("/admin", admin.handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}