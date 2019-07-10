package handlers

import (
	"log"
	"net/http"
	"encoding/json"
	"idcyberark/counter"
	"github.com/gorilla/mux"
)

type ResponseIdCyberark struct {
    Id   int `json:"id"`
}

func idcyberark(c *counter.SafeCounter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		idCyberArk := c.Inc(params["jenkins"])

		log.Printf("Id attribué : %d à %s", idCyberArk, params["jenkins"])

		responseBody := ResponseIdCyberark{idCyberArk}

		data, err := json.Marshal(responseBody)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}
