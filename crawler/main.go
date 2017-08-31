package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/web-crawler/scraper"
)

const (
	fileTemplate = "dump/gpdump"
	fileExt      = "csv"
)

func main() {
	n1 := 91380 // min for now
	n2 := 9999999

	taskCh := make(chan int)
	nWorkers := 100
	var wg sync.WaitGroup
	wg.Add(nWorkers)

	for i := 0; i < nWorkers; i++ {
		go worker(taskCh, i, &wg)
	}
	for i := n1; i <= n2; i++ {
		taskCh <- i
	}
	// tasks sent
	close(taskCh)
	// wait for workers to finist
	wg.Wait()
}

func worker(taskCh chan int, id int, wg *sync.WaitGroup) {
	defer wg.Done()

	filename := fmt.Sprintf("%s%d.%s", fileTemplate, id, fileExt)

	// create file for dumping
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	for task := range taskCh {
		// get data from page
		link, date, title, category, body, err := scraper.ScrapPage(task)
		if err == nil {
			lineToFile := constructLine(link, date, title, category, body)
			if _, err = f.WriteString(lineToFile); err != nil {
				panic(err)
			}
			if _, err = f.WriteString("\r\n"); err != nil {
				panic(err)
			}

		} else {
			fmt.Printf("Error for page %d, %v\n", id, err)
		}
	}
}

func constructLine(link, date, title, category, body string) string {
	return fmt.Sprintf(`"%s"|"%s"|"%s"|"%s"|"%s"`, link, date, title, category, body)
}
