/*
Solution to Exercise: Web Crawler https://tour.golang.org/concurrency/10
*/

package main

import (
	"fmt"
	"sync"
)

// プロセス単位排他
var processMutex = new(sync.Mutex)

type UrlInfos struct {
	// URL文字列を KEY にして、対象URL の情報を管理する map
	urlMap map[string]UrlInfo
}

type UrlInfo struct {
	// crawl 済みフラグ
	IsCrawled bool
	// URL 単位排他
  UrlMutex *sync.Mutex
}

func (u UrlInfos) Lock(url string) {
	processMutex.Lock()
	defer processMutex.Unlock()

	urlInfo, exist := u.urlMap[url]

	if exist == false {
		urlInfo = UrlInfo{}
		urlInfo.UrlMutex = new(sync.Mutex)
		u.urlMap[url] = urlInfo
	}

	urlInfo.UrlMutex.Lock()
}

func (u UrlInfos) Unlock(url string) {
	processMutex.Lock()
	defer processMutex.Unlock()

	urlInfo := u.urlMap[url]

	urlInfo.UrlMutex.Unlock()
}

func (u UrlInfos) SetCrawled(url string) {
	processMutex.Lock()
	defer processMutex.Unlock()

	urlInfo := u.urlMap[url]

	urlInfo.IsCrawled = true
	u.urlMap[url] = urlInfo
}

func (u UrlInfos) IsCrawled(url string) bool {
	processMutex.Lock()
	defer processMutex.Unlock()

	urlInfo := u.urlMap[url]

	return urlInfo.IsCrawled
}

var urlInfos = UrlInfos{make(map[string]UrlInfo)}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	wg := sync.WaitGroup{}

	StartCrawlWorker(url, depth, fetcher, &wg)

	// 全　CrawlWorker の終了を待つ
	wg.Wait()
}

func StartCrawlWorker(url string, depth int, fetcher Fetcher, wg *sync.WaitGroup) {
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}

	// URL 単位の排他ロック
	urlInfos.Lock(url)
	defer urlInfos.Unlock(url)

	if urlInfos.IsCrawled(url) {
		// crawl 済み
		return
	}

	// sync.WaitGroup をカウントアップ
	// goroutine は実行タイミングが任意のため、goroutine 内部に記述するとカウントアップ前に
	// wg.Wait() を抜けて、プロセスが終了してしまうので実行前に記述する必要あり
	wg.Add(1)

	go CrawlWorker(url, depth, fetcher, wg)
}

func CrawlWorker(url string, depth int, fetcher Fetcher, wg *sync.WaitGroup) {
	defer wg.Done()

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)

	// body 取得済みフラグセット
	urlInfos.SetCrawled(url)

	for _, u := range urls {
		StartCrawlWorker(u, depth-1, fetcher, wg)
	}
	return
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}

	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
