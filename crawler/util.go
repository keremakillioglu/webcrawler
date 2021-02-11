package crawler

import (
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

const (
	// PageExtension is the file extension for the downloaded pages
	PageExtension = ".html"
	// PageDirIndex is the file name of the index file for every dir
	PageDirIndex = "index" + PageExtension
)

// randomString is for creating random user agents
func randomString() string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, rand.Intn(10)+10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// getPageFilePath returns a filename for a URL that represents a page.
func getPageFilePath(url *url.URL) string {

	fileName := url.Path
	if len(url.Path) >= 1 {
		fileName = url.Path[1:]
	}

	// root of domain will be index.html
	if fileName == "" || fileName == "/" {
		fileName = PageDirIndex

		// directory index will be index.html in the directory
	} else if fileName[len(fileName)-1] == '/' {
		fileName += PageDirIndex

	} else {
		ext := filepath.Ext(fileName)
		// if there is no file extension, add .html
		if ext == "" {
			fileName += PageExtension
		} else if ext != PageExtension {
			// replace other extensions with .html
			fileName = fileName[:len(fileName)-len(ext)] + PageExtension
		}
	}

	return fileName
}

// writeFile creates the visited HTML file, and maps it to its directory
func writeFile(url *url.URL, filePath string, res *http.Response) error {

	// directory is initiated with url host name
	dir := filepath.Dir(url.Host + "/" + filePath)

	// append to root directory
	if len(dir) < len(url.Host) {
		dir = filepath.Join(".", url.Host, dir)
	}

	// mkdir
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}

	// if no directory
	if !strings.Contains(filePath, "/") {
		filePath = url.Host + "/" + filePath
	}

	// map existing directories to filepath
	pos := strings.LastIndexAny(filePath, "/")
	if pos != -1 {
		firstPart := filePath[:pos]
		if strings.Contains(dir, firstPart) {
			filePath = dir + "/" + filePath[pos+1:]
		}
	}

	// create output file
	outFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// copy data from HTTP response to file
	_, err = io.Copy(outFile, res.Body)
	if err != nil {
		return err
	}

	return err
}
