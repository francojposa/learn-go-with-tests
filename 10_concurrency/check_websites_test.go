package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	if url == "http://exampl.ecom" {
		return false
	}
	return true
}

func TestWebsitesChecker(t *testing.T) {

	urls := []string{
		"https://google.com",
		"http://twitter.com",
		"http://exampl.ecom",
	}

	want := map[string]bool{
		"https://google.com": true,
		"http://twitter.com": true,
		"http://exampl.ecom": false,
	}

	got := CheckWebsites(mockWebsiteChecker, urls)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("\nwant: %v, got: %v", want, got)
	}
}

func slowWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkWebsitesChecker(b *testing.B) {
	urls := make([]string, 100)
	for i := range urls {
		urls[i] = "stub url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowWebsiteChecker, urls)
	}
}
