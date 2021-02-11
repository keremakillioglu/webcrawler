# Webcrawler
A recursirve web-crawler built with Golang

## Usage
- go run main.go
- Enter URL to Crawl: http://keremakillioglu.github.io

## Web-Crawler Implementation

### Dependencies
* **Colly:** Web scraping
* **Goquery:** HTML operations
* **Logrus:** Logging 

## Input
Web crawler takes URL as a command line input. If input is not validated, it will be asked again.
Correct formats: "http://keremakillioglu.github.io", "http://people.sabanciuniv.edu/ysaygin/", "https://www.levels.fyi/"

## Allowed Domains
Crawler does not visit any pages with different domains other than input.
Domain restriction is handled through Colly Collector configurations.

## Program Flow
* Crawler can visit multiple pages. The number of web-crawling goroutines can be set. (Default set as: 4)
* Crawler calls getdata() to download the corresponding url
* File directories are mapped with getPageFilePath()
* After basic checks, writeFile() writes data to desired directory

#### Limitations
##### File Types
Only HTML files are supported for download.
PDF, PPT, XML or other file types are downloaded as html files with a margin of error.
Assets like CSS files cannot be downloaded.

##### File Types
Destination directory is by default set as project directory

##### Number of Hosts
When crawling with a single host it may fail with response code:429.
It is because crawler is using more threads than the server allows.
Reduce parallelism in that case.

##### Further Development
Instead of I/O, Cobra package can be used for cmd operations.
Logs can be stored
