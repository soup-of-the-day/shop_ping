package watcher

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// Grab some html page and allow for scraping
type HtmlWatcher struct {
	url string
}

func CreateHtmlWatcher(url string) *HtmlWatcher {
	return &HtmlWatcher{
		url: url,
	}
}

// Determine if some string exists on a web page
func (hw *HtmlWatcher) Find(s string) (bool, error) {
	resp, err := http.Get(hw.url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	return strings.Contains(string(html), s), nil
}
