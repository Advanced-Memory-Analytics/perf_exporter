package perf

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/Advanced-Memory-Analytics/perf_exporter/util"
)

type Mem struct {
	Load MemLoads `json:"load,omitempty"`
}

type MemHeader struct {
	SampleCount int    `json:"sample_count,omitempty"`
	EventCount  int    `json:"event_count,omitempty"`
	Event       string `json:"event,omitempty"`
}

type MemLoads struct {
	Header MemHeader       `json:"header,omitempty"`
	Data   []MemStatsTable `json:"data,omitempty"`
}

type MemStatsTable struct {
	OverheadPCT   float64 `json:"percentage"`
	Samples       int     `json:"samples"`
	MemAccessType string  `json:"access"`
}

func MemCollector(collectorType string) (string, error) {
	switch collectorType {
	case "load":
		{
			perf := Perf{}

			err := os.WriteFile("mem_load.txt", []byte(util.GenerateRandMemLoadString()), os.ModeExclusive)
			if err != nil {
				return "", fmt.Errorf("failed to write to file %v", err)
			}

			text, err := os.ReadFile("mem_load.txt")
			if err != nil {
				return "", fmt.Errorf("failed to read file %v", err)
			}

			err = perf.ParseMemLoads(string(text))
			if err != nil {
				log.Debug().Msgf("failed to parse memory load file %v", err)
				return "", fmt.Errorf("failed to parse memory load file %v", err)
			}

			jsonData, err := json.Marshal(perf)
			if err != nil {
				return "", fmt.Errorf("failed to marshal memory load file results %v", err)
			}
			return string(jsonData), nil
		}
	case "store":
		{
			time.Sleep(2 * time.Second)
			perf := Perf{}

			text, err := os.ReadFile("mem_store.txt")
			if err != nil {
				return "", fmt.Errorf("failed to read file %v", err)
			}

			err = perf.ParseMemLoads(string(text))
			if err != nil {
				log.Debug().Msgf("failed to parse memory store file %v", err)
				return "", fmt.Errorf("failed to parse memory load file %v", err)
			}

			fmt.Printf("Perf Mem Stores: %v\n", perf.Mem.Load)

			jsonData, err := json.Marshal(perf)
			if err != nil {
				return "", fmt.Errorf("failed to marshal memory store file results %v", err)
			}
			return string(jsonData), nil
		}
	default:
		return "", fmt.Errorf("something went wrong")
	}
}