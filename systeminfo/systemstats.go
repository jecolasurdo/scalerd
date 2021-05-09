package systeminfo

import (
	"time"

	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/memory"
)

const cpuObservationWindowDuration time.Duration = 1000 * time.Millisecond

type SystemStats struct {
	CPUPercent    float64
	MemoryPercent float64
}

func MustGetSystemStats() SystemStats {
	memory, err := memory.Get()
	if err != nil {
		panic(err)
	}
	memoryPct := 100 * float64(memory.Used) / float64(memory.Total)

	before, err := cpu.Get()
	if err != nil {
		panic(err)
	}
	time.Sleep(cpuObservationWindowDuration)
	after, err := cpu.Get()
	if err != nil {
		panic(err)
	}
	cpuPct := (float64(after.User-before.User) + float64(after.System-before.System)) / float64(after.Total-before.Total) * 100

	return SystemStats{
		MemoryPercent: memoryPct,
		CPUPercent:    cpuPct,
	}
}
