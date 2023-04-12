package api

import (
	"EndGo/data"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func InfoOfCity(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	idCity := chi.URLParam(r, "id")

	cityFound := false
	for _, city := range data.ArrCities {
		if city.Id == idCity {
			output, err := json.Marshal(city)
			if err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				rw.Write([]byte(err.Error()))
				return
			}
			rw.Write(output)
			rw.Write([]byte("\n"))
			cityFound = true
			break
		}
	}
	if cityFound == false {
		rw.Write([]byte("Город не найден\n"))
		rw.WriteHeader(http.StatusNotFound)
	}
	return
}
