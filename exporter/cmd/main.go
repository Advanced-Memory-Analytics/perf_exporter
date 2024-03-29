package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"

	promCollector "github.com/Advanced-Memory-Analytics/perf_exporter/internal/prometheus_collector"
	"github.com/Advanced-Memory-Analytics/perf_exporter/util"
)

func init() {
	prometheus.Unregister(collectors.NewGoCollector())
	prometheus.MustRegister(promCollector.NewCollector())
}
func main() {
	config, err := util.LoadConfig()
	if err != nil {
		log.Fatal().Msgf("Failed to load config %v", err)
	}
	port := config.SERVER_ADDRESS

	log.Debug().Msgf("||Config|| server_addr: %v | kafka_addr: %v", config.SERVER_ADDRESS, config.KAFKA_ADDRESS)

	http.Handle("/metrics", promhttp.Handler())
	log.Logger.Fatal().Msgf("err: %v", http.ListenAndServe(fmt.Sprintf(":%s", port), nil))

}
