package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	if url == "http://examplec.om" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"https://duckduckgo.com",
		"http://twitter.com",
		"http://examplec.om",
	}

	want := map[string]bool{
		"https://duckduckgo.com": true,
		"http://twitter.com":     true,
		"http://examplec.om":     false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("\nwant: %v, got: %v", want, got)
	}
}
