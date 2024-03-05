package promCollector

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/Advanced-Memory-Analytics/perf_exporter/internal/perf"
)

// var ToPerf chan string = make(chan string)

var PromChannel chan string

type collector struct {
	memLoadSampleCount *prometheus.Desc
	memLoadEventCount  *prometheus.Desc
	memEventType       *prometheus.Desc

	memLoadLFBHitSamplePct   *prometheus.Desc
	memLoadLFBHitSampleCount *prometheus.Desc

	memLoadL1HitSampleCount *prometheus.Desc
	memLoadL1HitSamplePct   *prometheus.Desc

	memLoadL3HitSampleCount *prometheus.Desc
	memLoadL3HitSamplePct   *prometheus.Desc

	memLoadLDLRAMHitSampleCount *prometheus.Desc
	memLoadLDLRAMHitSamplePct   *prometheus.Desc

	memLoadL2HitSampleCount *prometheus.Desc
	memLoadL2HitSamplePct   *prometheus.Desc

	memLoadUncachedNAHitSampleCount *prometheus.Desc
	memLoadUncachedNAHitSamplePct   *prometheus.Desc

	memLoadIONAHitSampleCount *prometheus.Desc
	memLoadIONAHitSamplePct   *prometheus.Desc

	memLoadL3MissHitSampleCount *prometheus.Desc
	memLoadL3MissHitSamplePct   *prometheus.Desc
}

func NewCollector() *collector {
	return &collector{
		memLoadSampleCount: prometheus.NewDesc(
			"mem_load_sample_count",
			"Number of memory load samples.",
			[]string{"event"},
			nil,
		),
		memLoadEventCount: prometheus.NewDesc(
			"mem_load_event_count",
			"Number of memory load events.",
			[]string{"event"},
			nil,
		),
		memEventType: prometheus.NewDesc(
			"mem_event_type",
			"Type of memory event.",
			[]string{"event"},
			nil,
		),
		memLoadLFBHitSamplePct: prometheus.NewDesc(
			"mem_load_lfb_hit_sample_pct",
			"Percentage of memory loads accessing LFB or LFB hit.",
			[]string{"access"},
			nil,
		),
		memLoadLFBHitSampleCount: prometheus.NewDesc(
			"mem_load_lfb_hit_sample_count",
			"Number of memory loads accessing LFB or LFB hit.",
			[]string{"access"},
			nil,
		),
		memLoadL1HitSampleCount: prometheus.NewDesc(
			"mem_load_l1_hit_sample_count",
			"Number of memory loads accessing L1 or L1 hit.",
			[]string{"access"},
			nil,
		),
		memLoadL1HitSamplePct: prometheus.NewDesc(
			"mem_load_l1_hit_sample_pct",
			"Percentage of memory loads accessing L1 or L1 hit.",
			[]string{"access"},
			nil,
		),
		memLoadL3HitSampleCount: prometheus.NewDesc(
			"mem_load_l3_hit_sample_count",
			"Number of memory loads accessing L3 or L3 hit.",
			[]string{"access"},
			nil,
		),
		memLoadL3HitSamplePct: prometheus.NewDesc(
			"mem_load_l3_hit_sample_pct",
			"Percentage of memory loads accessing L3 or L3 hit.",
			[]string{"access"},
			nil,
		),
		memLoadLDLRAMHitSampleCount: prometheus.NewDesc(
			"mem_load_ldl_ram_hit_sample_count",
			"Number of memory loads accessing Ldl RAM or RAM hit.",
			[]string{"access"},
			nil,
		),
		memLoadLDLRAMHitSamplePct: prometheus.NewDesc(
			"mem_load_ldl_ram_hit_sample_pct",
			"Percentage of memory loads accessing Ldl RAM or RAM hit.",
			[]string{"access"},
			nil,
		),
		memLoadL2HitSampleCount: prometheus.NewDesc(
			"mem_load_l2_hit_sample_count",
			"Number of memory loads accessing L2 or L2 hit.",
			[]string{"access"},
			nil,
		),
		memLoadL2HitSamplePct: prometheus.NewDesc(
			"mem_load_l2_hit_sample_pct",
			"Percentage of memory loads accessing L2 or L2 hit.",
			[]string{"access"},
			nil,
		),
		memLoadUncachedNAHitSampleCount: prometheus.NewDesc(
			"mem_load_uncached_na_hit_sample_count",
			"Number of memory loads accessing Uncached or N/A hit.",
			[]string{"access"},
			nil,
		),
		memLoadUncachedNAHitSamplePct: prometheus.NewDesc(
			"mem_load_uncached_na_hit_sample_pct",
			"Percentage of memory loads accessing Uncached or N/A hit.",
			[]string{"access"},
			nil,
		),
		memLoadIONAHitSampleCount: prometheus.NewDesc(
			"mem_load_iona_hit_sample_count",
			"Number of memory loads accessing I/O or N/A hit.",
			[]string{"access"},
			nil,
		),
		memLoadIONAHitSamplePct: prometheus.NewDesc(
			"mem_load_iona_hit_sample_pct",
			"Percentage of memory loads accessing I/O or N/A hit.",
			[]string{"access"},
			nil,
		),
		memLoadL3MissHitSampleCount: prometheus.NewDesc(
			"mem_load_l3_miss_hit_sample_count",
			"Number of memory loads accessing L3 miss.",
			[]string{"access"},
			nil,
		),
		memLoadL3MissHitSamplePct: prometheus.NewDesc(
			"mem_load_l3_miss_hit_sample_pct",
			"Percentage of memory loads accessing L3 miss.",
			[]string{"access"},
			nil,
		),
	}
}

func (c *collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.memLoadSampleCount
	ch <- c.memLoadEventCount
	ch <- c.memEventType

	ch <- c.memLoadLFBHitSamplePct
	ch <- c.memLoadLFBHitSampleCount

	ch <- c.memLoadL1HitSampleCount
	ch <- c.memLoadL1HitSamplePct

	ch <- c.memLoadL3HitSampleCount
	ch <- c.memLoadL3HitSamplePct

	ch <- c.memLoadLDLRAMHitSampleCount
	ch <- c.memLoadLDLRAMHitSamplePct

	ch <- c.memLoadL2HitSampleCount
	ch <- c.memLoadL2HitSamplePct

	ch <- c.memLoadUncachedNAHitSampleCount
	ch <- c.memLoadUncachedNAHitSamplePct

	ch <- c.memLoadIONAHitSampleCount
	ch <- c.memLoadIONAHitSamplePct

	ch <- c.memLoadL3MissHitSampleCount
	ch <- c.memLoadL3MissHitSamplePct
}

type MemoryLoadData struct {
	Percentage float64 `json:"percentage"`
	Samples    int     `json:"samples"`
	Access     string  `json:"access"`
}

type MemoryLoadJSON struct {
	Mem struct {
		Load struct {
			Header struct {
				SampleCount int    `json:"sample_count"`
				EventCount  int    `json:"event_count"`
				Event       string `json:"event"`
			} `json:"header"`
			Data []MemoryLoadData `json:"data"`
		} `json:"load"`
	} `json:"mem"`
}

func (c *collector) Collect(ch chan<- prometheus.Metric) {
	// Parse the JSON data
	// jsonData := `{
	// 	"mem": {
	// 		"load": {
	// 			"header": {
	// 				"sample_count": 57000,
	// 				"event_count": 60000,
	// 				"event": "cpu/mem-loads,ldlat=30/P"
	// 			},
	// 			"data": [
	// 				{
	// 					"percentage": 21.34,
	// 					"samples": 12349,
	// 					"access": "LFB or LFB hit"
	// 				},
	// 				{
	// 					"percentage": 19.83,
	// 					"samples": 11476,
	// 					"access": "L1 or L1 hit"
	// 				},
	// 				{
	// 					"percentage": 17.11,
	// 					"samples": 9901,
	// 					"access": "L3 or L3 hit"
	// 				},
	// 				{
	// 					"percentage": 14.91,
	// 					"samples": 8628,
	// 					"access": "Ldl RAM or RAM hit"
	// 				},
	// 				{
	// 					"percentage": 10.09,
	// 					"samples": 5841,
	// 					"access": "L2 or L2 hit"
	// 				},
	// 				{
	// 					"percentage": 7.89,
	// 					"samples": 4568,
	// 					"access": "Uncached or N/A hit"
	// 				},
	// 				{
	// 					"percentage": 5.17,
	// 					"samples": 2993,
	// 					"access": "I/O or N/A hit"
	// 				},
	// 				{
	// 					"percentage": 3.66,
	// 					"samples": 2120,
	// 					"access": "L3 miss"
	// 				}
	// 			]
	// 		}
	// 	}
	// }`

	jsonData, err := perf.MemCollector("load")
	if err != nil {
		log.Fatalf("Failed to generate JSON data: %v", err)
	}

	fmt.Printf("TO PERF: %v\n", jsonData)
	var data MemoryLoadJSON
	err = json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON data: %v", err)
	}

	ch <- prometheus.MustNewConstMetric(c.memLoadSampleCount, prometheus.GaugeValue, float64(data.Mem.Load.Header.SampleCount), fmt.Sprintf("samples:%s", data.Mem.Load.Header.Event))
	ch <- prometheus.MustNewConstMetric(c.memLoadEventCount, prometheus.GaugeValue, float64(data.Mem.Load.Header.EventCount), fmt.Sprintf("events:%s", data.Mem.Load.Header.Event))

	for _, entry := range data.Mem.Load.Data {
		switch entry.Access {
		case "LFB or LFB hit":
			// Handle LFB or LFB hit case
			ch <- prometheus.MustNewConstMetric(c.memLoadLFBHitSampleCount, prometheus.GaugeValue, float64(entry.Samples), entry.Access)
			ch <- prometheus.MustNewConstMetric(c.memLoadLFBHitSamplePct, prometheus.GaugeValue, entry.Percentage, entry.Access)
		case "L1 or L1 hit":
			// Handle L1 or L1 hit case
			ch <- prometheus.MustNewConstMetric(c.memLoadL1HitSampleCount, prometheus.GaugeValue, float64(entry.Samples), entry.Access)
			ch <- prometheus.MustNewConstMetric(c.memLoadL1HitSamplePct, prometheus.GaugeValue, entry.Percentage, entry.Access)
		case "L3 or L3 hit":
			// Handle L3 or L3 hit case
			ch <- prometheus.MustNewConstMetric(c.memLoadL3HitSampleCount, prometheus.GaugeValue, float64(entry.Samples), entry.Access)
			ch <- prometheus.MustNewConstMetric(c.memLoadL3HitSamplePct, prometheus.GaugeValue, entry.Percentage, entry.Access)
		case "Ldl RAM or RAM hit":
			// Handle Ldl RAM or RAM hit case
			ch <- prometheus.MustNewConstMetric(c.memLoadLDLRAMHitSampleCount, prometheus.GaugeValue, float64(entry.Samples), entry.Access)
			ch <- prometheus.MustNewConstMetric(c.memLoadLDLRAMHitSamplePct, prometheus.GaugeValue, entry.Percentage, entry.Access)
		case "L2 or L2 hit":
			// Handle L2 or L2 hit case
			ch <- prometheus.MustNewConstMetric(c.memLoadL2HitSampleCount, prometheus.GaugeValue, float64(entry.Samples), entry.Access)
			ch <- prometheus.MustNewConstMetric(c.memLoadL2HitSamplePct, prometheus.GaugeValue, entry.Percentage, entry.Access)
		case "Uncached or N/A hit":
			// Handle Uncached or N/A hit case
			ch <- prometheus.MustNewConstMetric(c.memLoadUncachedNAHitSampleCount, prometheus.GaugeValue, float64(entry.Samples), entry.Access)
			ch <- prometheus.MustNewConstMetric(c.memLoadUncachedNAHitSamplePct, prometheus.GaugeValue, entry.Percentage, entry.Access)
		case "I/O or N/A hit":
			// Handle I/O or N/A hit case
			ch <- prometheus.MustNewConstMetric(c.memLoadIONAHitSampleCount, prometheus.GaugeValue, float64(entry.Samples), entry.Access)
			ch <- prometheus.MustNewConstMetric(c.memLoadIONAHitSamplePct, prometheus.GaugeValue, entry.Percentage, entry.Access)
		case "L3 miss":
			// Handle L3 miss case
			ch <- prometheus.MustNewConstMetric(c.memLoadL3MissHitSampleCount, prometheus.GaugeValue, float64(entry.Samples), entry.Access)
			ch <- prometheus.MustNewConstMetric(c.memLoadL3MissHitSamplePct, prometheus.GaugeValue, entry.Percentage, entry.Access)
		}
	}
}
