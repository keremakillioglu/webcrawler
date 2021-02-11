package crawler

import (
	"net/url"
	"time"

	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
)

// Crawl initiates web crawling operations, downloads the initial url, then check links inside original document recursively
func Crawl(fetchURL *url.URL) error {

	// download the initial url
	err := getdata(fetchURL.String())
	if err != nil {
		log.Error("Initial URL Download Error %v", err)
		return err
	}

	// colly configurations
	c := colly.NewCollector(
		colly.AllowedDomains(fetchURL.Hostname(), "www."+fetchURL.Host),
		colly.Async(true),
	)
	c.SetRequestTimeout(100 * time.Second)
	c.Limit(&colly.LimitRule{Parallelism: 4})
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", randomString())
		log.Info("Visiting ", r.URL.String())

	})

	// error logs
	c.OnError(func(r *colly.Response, err error) {
		log.Error("Request URL:", r.Request.URL, "failed with response code:", r.StatusCode, "\nError:", err)
	})

	// check links inside original document recursively
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {

		link := e.Attr("href")

		cURL, _ := url.ParseRequestURI(e.Request.AbsoluteURL(link))
		if cURL != nil {
			if cURL.Hostname() == fetchURL.Hostname() {
				err = getdata(e.Request.AbsoluteURL(link))
				if err != nil {
					log.Error("Download Error: ", err)
					return
				}
			}
			// visit link found on page on a new thread
			e.Request.Visit(e.Request.AbsoluteURL(link))
		}

	})

	// start web crawling
	c.Visit(fetchURL.String())

	// wait until threads are finished
	c.Wait()

	return err

}
