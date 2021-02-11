package crawler

import (
	"net/http"
	"net/url"
)

func getdata(fetchURL string) error {

	u, err := url.ParseRequestURI(fetchURL)
	if err != nil {
		return err
	}

	// e.g. if scheme is mailto, there is no error but cannot download
	if !(u.Scheme == "https" || u.Scheme == "http") {
		return nil
	}

	// make request
	response, err := http.Get(fetchURL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	fp := getPageFilePath(u)

	err = writeFile(u, fp, response)
	if err != nil {
		return err
	}

	return err
}
