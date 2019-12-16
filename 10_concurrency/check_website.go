package concurrency

import "net/http"

func HttpWebsiteChecker(url string) bool {
	response, err := http.Head(url)
	if err != nil {
		return false
	}
	if response.StatusCode != 200 {
		return false
	}
	return true
}
