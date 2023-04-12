package data

import (
	"github.com/gocarina/gocsv"
	"os"
)

type City struct {
	Id         string `csv:"id" json:"id"`
	Name       string `csv:"name" json:"name"`
	Region     string `csv:"region" json:"region"`
	District   string `csv:"district" json:"district"`
	Population string `csv:"population" json:"population"`
	Foundation string `csv:"foundation" json:"foundation"`
}

var ArrCities []*City

func UnmarshalCSV() {
	file, err := os.Open("cities.csv")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	if err := gocsv.UnmarshalFile(file, &ArrCities); err != nil {
		panic(err)
	}
}

func MarshalCSV() {
	file, err := os.OpenFile("cities.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	if _, err := file.Seek(0, 0); err != nil {
		panic(err)
	}

	err = gocsv.MarshalFile(&ArrCities, file)
	if err != nil {
		panic(err)
	}
}
