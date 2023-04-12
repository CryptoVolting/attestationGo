package main

import (
	"EndGo/api"
	"EndGo/data"
	"context"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	server := &http.Server{Addr: "localhost:9000", Handler: service()}
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	go func() {
		<-sig

		shutdownCtx, _ := context.WithTimeout(serverCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out.. Принудительный выход.")
			}
		}()

		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}
		serverStopCtx()
	}()

	data.UnmarshalCSV() // Запись информации о городах из cities.csv в массив ArrCities

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
	data.MarshalCSV() // Перезапись информации о городах в файл cities.csv из массива ArrCities
	log.Println("Сервер завершает свою работу.")
	<-serverCtx.Done()
}

func service() http.Handler {
	rout := chi.NewRouter()
	rout.Get("/info/{id}", api.InfoOfCity)
	rout.Post("/new", api.NewCity)
	rout.Delete("/delete", api.DeleteCity)
	rout.Put("/update/{id}", api.UpdatePopulation)
	rout.Post("/search", api.CitySearch)
	return rout
}
