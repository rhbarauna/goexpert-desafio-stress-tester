package stresstester

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/rhbarauna/goexpert-desafio-stress-tester/internal/report"
)

func RunTester(url string, requests int, concurrency int) {
	log.Printf("Starting Stress Tester... \n\nParams:\n  url\t\t=> %s \n  requests\t=> %d \n  concurrency\t=> %d\n\n", url, requests, concurrency)

	report := report.NewReport(url, requests, concurrency)

	if concurrency > requests {
		concurrency = requests
	}

	total_loops := requests / concurrency
	remainder := requests % concurrency

	if remainder > 0 {
		total_loops += 1
	}

	loops := make([]int, total_loops)

	for i := 0; i < total_loops; i++ {
		if i == total_loops-1 && remainder > 0 {
			loops[i] = remainder
			continue
		}
		loops[i] = concurrency
	}

	log.Printf("Executing")
	wg := &sync.WaitGroup{}

	for _, loop := range loops {
		wg.Add(loop)
		log.Printf("Dispatching %d\n", loop)
		for j := 0; j < loop; j++ {
			go executeRequest(url, wg, report)
		}
		wg.Wait()
		fmt.Println()
	}
	fmt.Printf(" OK\n")
	log.Println("Done")
	report.PrintResults()
}

var httpClient = http.Client{}

func executeRequest(url string, wg *sync.WaitGroup, report *report.Report) {
	fmt.Printf(".")
	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		wg.Done()
		return
	}

	request.Close = true
	response, _ := httpClient.Do(request)
	defer response.Body.Close()
	report.RegisterResponseStatus(response.StatusCode)
	wg.Done()
}
