package main

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"

	"github.com/Advanced-Memory-Analytics/perf_exporter/perf"
	"github.com/Advanced-Memory-Analytics/perf_exporter/util"
)

func main() {
	perf := perf.Perf{}

	text, err := os.ReadFile("mem_load.txt")
	if err != nil {
		log.Fatal().Msgf("Failed to read file %v", err)
	}
	// fmt.Println(string(text)) // Print the content as a string

	err = perf.ParseMemLoads(string(text))
	if err != nil {
		log.Fatal().Msgf("Failed to parse: %v", err)
	}

	fmt.Printf("Perf Mem Loads: %v\n", perf.Mem.Load)

	config, err := util.LoadConfig()
	if err != nil {
		log.Fatal().Msgf("Failed to load config %v", err)
	}

	log.Debug().Msgf("||Config|| server_addr: %v | kafka_addr: %v", config.SERVER_ADDRESS, config.KAFKA_ADDRESS)

	// http.Handle("/metrics", promhttp.Handler())
	// log.Debug().Msgf("Server: perf_exporter started at %v.", config.SERVER_ADDRESS)

	// http.ListenAndServe(":"+config.SERVER_ADDRESS, nil)
}
