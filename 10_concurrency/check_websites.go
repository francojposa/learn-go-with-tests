package concurrency

type WebsiteChecker func(string) bool

type websiteResult struct {
	url   string
	check bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan websiteResult)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- websiteResult{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <-resultChannel
		results[result.url] = result.check
	}

	return results
}
