package api

import (
	"EndGo/data"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
)

type UpdatePopulationStruct struct {
	Population string `json:"new_population"`
}

func UpdatePopulation(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	cityId := chi.URLParam(r, "id")
	cityFound := false
	for i, city := range data.ArrCities {
		if city.Id == cityId {
			content, err := io.ReadAll(r.Body)
			defer r.Body.Close()
			if err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				rw.Write([]byte(err.Error()))
				return
			}

			var newPop UpdatePopulationStruct

			if err = json.Unmarshal(content, &newPop); err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				rw.Write([]byte(err.Error()))
				return
			}
			data.ArrCities[i].Population = newPop.Population
			cityFound = true
			rw.Write([]byte("Численность населения обновлена у города с Id:" + cityId + "\n"))
			break
		}

	}
	if cityFound == false {
		rw.Write([]byte("Город не найден\n"))
		rw.WriteHeader(http.StatusNotFound)
	}
	return
}
