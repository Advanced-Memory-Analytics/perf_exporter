package perf

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/Advanced-Memory-Analytics/perf_exporter/util"
)

func (p *Perf) ParseMemLoads(text string) error {

	// Grab first two lines of text
	headerLines := strings.SplitN(text, "\n", 2)
	if len(headerLines) < 2 {
		return fmt.Errorf("invalid format: missing header")
	}

	// Parse the Header line
	header := MemHeader{}
	load := MemLoads{}

	re := regexp.MustCompile(`^Samples: (?P<samples>\d+K) of event '(?P<event>[^']+)'(?P<extra>.+)?, Event count \(approx.\): (?P<event_count>\d+)\s*$`)
	match := re.FindStringSubmatch(headerLines[0])
	if match == nil {
		return fmt.Errorf("unable to match %s", headerLines[0])
	}

	sampleCount, err := util.ConvertNumberFormat(match[1])
	if err != nil {
		return fmt.Errorf("sample count in invalid format: %v", err)
	}

	header.SampleCount = sampleCount
	header.Event = match[2]
	header.EventCount, _ = strconv.Atoi(match[4])

	load.Header = header

	var data []MemStatsTable
	lines := strings.Split(headerLines[1], "\n")[1:] // Skipping first header line
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) < 3 {
			return fmt.Errorf("invalid format: MemStatsTable line has less than 3 fields: %s", line)
		}

		// Overhead Percentage
		percentage, err := strconv.ParseFloat(fields[0][:len(fields[0])-1], 64) // Exclude "%" symbol
		if err != nil {
			return fmt.Errorf("invalid format: parsing percentage: %w", err)
		}

		// Number of Samples
		samples, err := strconv.Atoi(fields[1])
		if err != nil {
			return fmt.Errorf("invalid format: parsing samples: %w", err)
		}

		// Memory Access
		memAccessType := strings.Join(fields[2:], " ")

		// Add to the data
		data = append(data, MemStatsTable{
			OverheadPCT:   percentage,
			Samples:       samples,
			MemAccessType: memAccessType,
		})
	}

	load.Data = data
	p.Mem.Load = load
	return nil
}
