package main

import (
	"fmt"
	"strings"

	"github.com/jecolasurdo/scalerd/systeminfo"
)

const (
	sampleBufferSize           int     = 100
	thresholdPercentile        float64 = 90
	cpuUtilizationThreshold    float64 = 90
	memoryUtilizationThreshold float64 = 80
)

func main() {
	systemStatHistory := systeminfo.NewSystemStatHistory(sampleBufferSize)
	for {
		systemStats := systeminfo.MustGetSystemStats()
		systemStatHistory.Update(systemStats)

		memoryPTile := systemStatHistory.MustGetMemoryPercentile(thresholdPercentile)
		cpuPTile := systemStatHistory.MustGetCPUPercenile(thresholdPercentile)

		// Process stats
		fmt.Println(strings.Repeat("-", 80))
		if memoryPTile <= memoryUtilizationThreshold && cpuPTile <= cpuUtilizationThreshold {
			fmt.Println("Below thresholds!")
			fmt.Printf("memory ptile: %v\n", memoryPTile)
			fmt.Printf("cpu ptile: %v\n", cpuPTile)
		}

	}

}
