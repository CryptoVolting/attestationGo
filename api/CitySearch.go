package api

import (
	"EndGo/data"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func CitySearch(rw http.ResponseWriter, r *http.Request) {
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

	if c.District != "" {
		for _, city := range data.ArrCities {
			if city.District == c.District {
				output, err := json.Marshal(city)
				if err != nil {
					rw.WriteHeader(http.StatusInternalServerError)
					rw.Write([]byte(err.Error()))
					return
				}
				rw.Write(output)
				rw.Write([]byte("\n"))
			}
		}
	}

	if c.Region != "" {
		for _, city := range data.ArrCities {
			if city.Region == c.Region {
				output, err := json.Marshal(city)
				if err != nil {
					rw.WriteHeader(http.StatusInternalServerError)
					rw.Write([]byte(err.Error()))
					return
				}
				rw.Write(output)
				rw.Write([]byte("\n"))
			}
		}
	}

	if c.Foundation != "" {
		minMaxFond := strings.Fields(c.Foundation)
		for _, city := range data.ArrCities {
			cityFond, _ := strconv.Atoi(city.Foundation)
			minFond, _ := strconv.Atoi(minMaxFond[0])
			maxFond, _ := strconv.Atoi(minMaxFond[1])
			if cityFond > minFond && cityFond < maxFond {
				output, err := json.Marshal(city)
				if err != nil {
					rw.WriteHeader(http.StatusInternalServerError)
					rw.Write([]byte(err.Error()))
					return
				}
				rw.Write(output)
				rw.Write([]byte("\n"))
			}
		}
	}

	if c.Population != "" {
		minMaxPop := strings.Fields(c.Population)
		for _, city := range data.ArrCities {
			cityPop, _ := strconv.Atoi(city.Population)
			minPop, _ := strconv.Atoi(minMaxPop[0])
			maxPop, _ := strconv.Atoi(minMaxPop[1])
			if cityPop > minPop && cityPop < maxPop {
				output, err := json.Marshal(city)
				if err != nil {
					rw.WriteHeader(http.StatusInternalServerError)
					rw.Write([]byte(err.Error()))
					return
				}
				rw.Write(output)
				rw.Write([]byte("\n"))
			}
		}
	}
	return
}
