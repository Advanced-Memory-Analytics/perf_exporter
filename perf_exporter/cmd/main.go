package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"

	"github.com/Advanced-Memory-Analytics/perf_exporter/util"
)

func main() {
	config, err := util.LoadConfig()
	if err != nil {
		log.Fatal().Msgf("Failed to load config %v", err)
	}

	http.Handle("/metrics", promhttp.Handler())
	log.Debug().Msgf("Server: perf_exporter started at %v.", config.SERVER_ADDRESS)

	http.ListenAndServe(":"+config.SERVER_ADDRESS, nil)
}
