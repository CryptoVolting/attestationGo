package api

import (
	"EndGo/data"
	"encoding/json"
	"io"
	"net/http"
)

type CityIdDeleteStruct struct {
	TargetID string `json:"target_id"`
}

func DeleteCity(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	content, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}
	var del CityIdDeleteStruct

	if err = json.Unmarshal(content, &del); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}
	//Удаление города из массива
	for i, city := range data.ArrCities {
		if city.Id == del.TargetID {
			data.ArrCities = append(data.ArrCities[:i], data.ArrCities[i+1:]...)
		}
	}
	rw.Write([]byte("Удален город с Id:" + del.TargetID + "\n"))
	return
}
