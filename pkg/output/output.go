package output

import (
	"fmt"
	"sync"
	"time"

	"github.com/divslingerx/sitewatcher/pkg/linkchecker"
)

func Generate(urls []string) string {
	ch := make(chan linkchecker.URLStatus)
	wg := new(sync.WaitGroup)

	for _, url := range urls {
		wg.Add(1)
		go linkchecker.CheckURL(url, wg, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	goodURLs := 0
	badURLs := make([]string, 0)
	for urlStatus := range ch {
		if urlStatus.Status == "🟢" {
			goodURLs++
		} else {
			badURLs = append(badURLs, urlStatus.URL)
		}
	}

	if goodURLs == len(urls) {
		return "🟢 All Sites Operational - " + time.Now().Format(time.Kitchen)
	} else {
		output := fmt.Sprintf("🟢 %d/%d Sites Operational - %s\n", goodURLs, len(urls), time.Now().Format(time.Kitchen))
		for _, url := range badURLs {
			output += fmt.Sprintf("🔴 WEBSITE DOWN: %s - %s\n", url, time.Now().Format(time.Kitchen))
		}
		return output
	}
}
