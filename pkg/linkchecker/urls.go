package linkchecker

import (
	"net/http"
	"sync"
)

type URLStatus struct {
	URL    string
	Status string
}

func CheckURL(url string, wg *sync.WaitGroup, ch chan<- URLStatus) {
	defer wg.Done()

	resp, err := http.Head(url)
	status := "🟢" // Green circle emoji for good URL
	if err != nil {
		status = "🔴" // Red circle emoji for bad URL
	} else {
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			status = "🔴" // Red circle emoji for bad URL
		}
	}

	ch <- URLStatus{URL: url, Status: status}
}
