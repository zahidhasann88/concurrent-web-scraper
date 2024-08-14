package scraper

// Result holds the result of a URL fetch operation
type Result struct {
	URL      string
	Duration float64
	Size     int
	Error    error
}
