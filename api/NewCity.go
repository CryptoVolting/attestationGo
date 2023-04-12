package api

import (
	"EndGo/data"
	"encoding/json"
	"io"
	"net/http"
)

func NewCity(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	content, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}
	var c *data.City

	if err = json.Unmarshal(content, &c); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}
	data.ArrCities = append(data.ArrCities, c)
	rw.Write([]byte("Город добавлен\n"))
	return
}
