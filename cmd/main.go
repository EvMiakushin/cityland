package main

import (
	"cityland/internal/control"
	"cityland/internal/depos"
	"cityland/internal/httpserv"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

const (
	Url    = "8081"
	Filer  = "cities.csv"
	Filer2 = "cities2.csv"
)

func main() {
	cb := depos.NewCityBase(Filer)
	story := depos.NewCityStor(cb)
	control := control.NewCityHand(story)
	handlers := httpserv.NewHandler(control)
	server := new(httpserv.Server)
	go func() {
		server.Start(Url, handlers.Https())
	}()
	fmt.Println("Server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	<-quit

	fmt.Println("Saving CSV file...")
	cb.SaveCityBase(Filer2)
	fmt.Println("Сities.csv file saved")
}
