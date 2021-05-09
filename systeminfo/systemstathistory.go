package systeminfo

import (
	"github.com/jecolasurdo/scalerd/buffers"
	"github.com/montanaflynn/stats"
)

type SystemStatHistory struct {
	cpuHistory    *buffers.CappedFloats
	memoryHistory *buffers.CappedFloats
}

func NewSystemStatHistory(sampleBufferSize int) SystemStatHistory {
	return SystemStatHistory{
		cpuHistory:    buffers.NewCappedFloats(sampleBufferSize),
		memoryHistory: buffers.NewCappedFloats(sampleBufferSize),
	}
}

func (s *SystemStatHistory) mustBeProperylInitialized() {
	if s.cpuHistory == nil || s.memoryHistory == nil {
		panic("SystemStatHistory must be initialized with NewSystemStatHistory")
	}
}

func (s *SystemStatHistory) Update(systemStats SystemStats) {
	s.mustBeProperylInitialized()
	s.cpuHistory.Push(systemStats.CPUPercent)
	s.memoryHistory.Push(systemStats.MemoryPercent)
}

func (s *SystemStatHistory) MustGetCPUPercenile(p float64) float64 {
	ptile, err := stats.Float64Data(s.cpuHistory.Data()).Percentile(p)
	if err != nil {
		panic(err)
	}
	return ptile
}

func (s *SystemStatHistory) MustGetMemoryPercentile(p float64) float64 {
	ptile, err := stats.Float64Data(s.memoryHistory.Data()).Percentile(p)
	if err != nil {
		panic(err)
	}
	return ptile
}
