package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
}

func main() {

	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "dev", "Environment (dev|staging|prod)")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheck)

	addr := fmt.Sprintf(":%d", cfg.port)

	logger.Printf("Starting %s server on %s", cfg.env, addr)

	err := http.ListenAndServe(addr, mux)
	if err != nil {
		fmt.Println(err)
	}
}
