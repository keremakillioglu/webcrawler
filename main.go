package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"runtime"

	"github.com/keremakillioglu/webcrawler/crawler"

	log "github.com/sirupsen/logrus"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	//taking input
	var err error = nil
	var fetchURL *url.URL
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter URL to Crawl: ")
		scanner.Scan()
		uinput := scanner.Text()
		fetchURL, err = url.ParseRequestURI(uinput)

		if err != nil {
			log.Error("Invalid URL Input")
		} else {
			break
		}
	}

	log.Info("Attempting to crawl: ", fetchURL.String())
	err = crawler.Crawl(fetchURL)
	if err != nil {
		log.Error("Scraping Error: %v", err)
	}

}
