package watcher

// Scrapes the contents of a web page
type Watcher interface {
	// Determine if a string exists in a page or not
	Find(s string) (bool, error)
}
