package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	_"os"
	"time"

	"github.com/Nisidhkr/olx-api/internal/config"
	// "github.com/joho/godotenv"
)

func main() {

	cfg:=config.MustLoad()


	

	mux := http.NewServeMux() //creting my own multiplexer so that we can add our own handlers to it and also we can add pprof handlers to it as pprof is internal package and we cannot add handlers to default mux

	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		//explicitly setting status code of our own
		w.WriteHeader(http.StatusOK) //oder of header and write matters, if we write before setting header then it will not work because write will set the status code to 200 by default and then we cannot change it after that

		w.Write([]byte(`{"status":"ok"}`))
	})
	fmt.Print("server running on :", "http://localhost:8090")

	srv := http.Server{
		Addr:         ":" +cfg.Port ,
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("server failed %v", err)
	}
}
