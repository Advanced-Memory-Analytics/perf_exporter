package main

import (
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/Advanced-Memory-Analytics/perf_exporter/internal/perf"
	"github.com/Advanced-Memory-Analytics/perf_exporter/util"
)

func main() {
	fmt.Println(util.GenerateRandMemLoadString())
	perf.MemCollector("load")

	config, err := util.LoadConfig()
	if err != nil {
		log.Fatal().Msgf("Failed to load config %v", err)
	}

	log.Debug().Msgf("||Config|| server_addr: %v | kafka_addr: %v", config.SERVER_ADDRESS, config.KAFKA_ADDRESS)
}
