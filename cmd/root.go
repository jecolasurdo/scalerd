package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/jecolasurdo/pacer"
	"github.com/jecolasurdo/scalerd/systeminfo"
)

const (
	samplingRate               float64       = 5
	samplingRateBasis          time.Duration = time.Second
	sampleBufferSize           int           = 100
	thresholdPercentile        float64       = 90
	cpuUtilizationThreshold    float64       = 90
	memoryUtilizationThreshold float64       = 90
)

func main() {
	pace := pacer.SetUniformPace(samplingRate, samplingRate, samplingRateBasis)

	systemStatHistory := systeminfo.NewSystemStatHistory(sampleBufferSize)
	for {
		systemStats := systeminfo.MustGetSystemStats()
		systemStatHistory.Update(systemStats)

		memoryPTile := systemStatHistory.MustGetMemoryPercentile(thresholdPercentile)
		cpuPTile := systemStatHistory.MustGetCPUPercenile(thresholdPercentile)

		fmt.Println(strings.Repeat("-", 80))
		if memoryPTile <= memoryUtilizationThreshold && cpuPTile <= cpuUtilizationThreshold {
			fmt.Println("Below thresholds!")
			fmt.Printf("memory ptile: %v\n", memoryPTile)
			fmt.Printf("cpu ptile: %v\n", cpuPTile)
		}
		pace.Wait()
	}

}
