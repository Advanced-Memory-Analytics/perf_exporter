package perf

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
