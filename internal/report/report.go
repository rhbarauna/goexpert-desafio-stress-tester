package report

import (
	"fmt"
	"sync"
	"time"
)

type Report struct {
	URL                   string
	ConfiguredRequests    int
	ConfiguredConcurrency int
	StartTime             time.Time
	TotalRequest          int
	ResponseStatuses      map[int]int
	mutex                 sync.Mutex
}

func NewReport(url string, requests int, concurrency int) *Report {
	report := &Report{}
	report.URL = url
	report.ConfiguredRequests = requests
	report.ConfiguredConcurrency = concurrency
	report.ResponseStatuses = map[int]int{}
	report.StartTime = time.Now()
	return report
}

func (r *Report) RegisterResponseStatus(statusCode int) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	currentCount, exists := r.ResponseStatuses[statusCode]

	if !exists {
		currentCount = 0
	}
	r.ResponseStatuses[statusCode] = currentCount + 1
}

func (r *Report) getDuration() string {
	d := time.Since(r.StartTime)

	if d < time.Second {
		return fmt.Sprintf("%dms", d.Milliseconds())
	}
	if d < time.Minute {
		return fmt.Sprintf("%.2f segundos", d.Seconds())
	}
	if d < time.Hour {
		return fmt.Sprintf("%d minutos, %.2f segundos", int(d.Seconds())/60, d.Seconds()-float64(int(d.Seconds())/60)*60)
	}

	return fmt.Sprintf("%d horas, %d minutos, %.2f segundos", int(d.Hours()), int(d.Minutes())%60, d.Seconds()-float64(int(d.Seconds())/60)*60)
}

func (r *Report) getTotalRequests() int {
	t := 0

	for _, v := range r.ResponseStatuses {
		t += v
	}

	return t
}

func (r *Report) PrintResults() {
	requests_completed := r.getTotalRequests()
	fmt.Println()

	fmt.Printf("\nTotal time: \t\t%s", r.getDuration())
	fmt.Printf("\nFailed requests: \t%d", r.ConfiguredRequests-requests_completed)
	fmt.Printf("\nCompleted Requests: \t%d", requests_completed)

	fmt.Printf("\n\nHTTP STATUS\tTOTAL")
	for key, value := range r.ResponseStatuses {
		fmt.Printf("\n%d\t\t%d", key, value)
	}

	fmt.Println()
}
