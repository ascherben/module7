package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
)

// URLs of Wikipedia pages to scrape
var urls = []string{
	"https://en.wikipedia.org/wiki/Robotics",
	"https://en.wikipedia.org/wiki/Robot",
	"https://en.wikipedia.org/wiki/Reinforcement_learning",
	"https://en.wikipedia.org/wiki/Robot_Operating_System",
	"https://en.wikipedia.org/wiki/Intelligent_agent",
	"https://en.wikipedia.org/wiki/Software_agent",
	"https://en.wikipedia.org/wiki/Robotic_process_automation",
	"https://en.wikipedia.org/wiki/Chatbot",
	"https://en.wikipedia.org/wiki/Applications_of_artificial_intelligence",
	"https://en.wikipedia.org/wiki/Android_(robot)",
}

// Page represents one Wikipedia page
type Page struct {
	URL     string `json:"url"`
	Content string `json:"text"`
	Time    string `json:"timestamp"`
}

// visits the Wikipedia pages and saves the results to a JSON Lines file
func Scrape(output string) error {
	file, err := os.Create(output)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
		colly.Async(true),
	)

	err = c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 2,
		Delay:       time.Second,
	})
	if err != nil {
		return err
	}

	// sends scraped pages from the crawler to the writer
	pages := make(chan Page)

	var writerErr error
	var writerWG sync.WaitGroup

	// Encodes pages to JSON and writes them to the file.
	writerWG.Add(1)
	go func() {
		defer writerWG.Done()

		for page := range pages {
			if writerErr != nil {
				continue
			}

			if err := encoder.Encode(page); err != nil {
				writerErr = err
			}
		}
	}()

	// Mutex and variable to capture the first error that occurs during scraping
	var errorLock sync.Mutex
	var scrapeErr error

	// capture the first scrape error
	saveError := func(err error) {
		errorLock.Lock()
		defer errorLock.Unlock()

		if scrapeErr == nil {
			scrapeErr = err
		}
	}

	// Logs each URL visited
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL.String())
	})

	// Extracts and cleans text from the main content of the page
	c.OnHTML("div.mw-parser-output", func(e *colly.HTMLElement) {
		var lines []string

		// Extract text from each paragraph and clean it
		e.ForEach("p", func(_ int, el *colly.HTMLElement) {
			content := clean(el.Text)

			if content == "" {
				return
			}

			lines = append(lines, content)
		})

		if len(lines) == 0 {
			return
		}

		page := Page{
			URL:     e.Request.URL.String(),
			Content: strings.Join(lines, "\n"),
			Time:    time.Now().Format(time.RFC3339),
		}

		pages <- page
	})

	// errors that occur during scraping
	c.OnError(func(r *colly.Response, err error) {
		saveError(fmt.Errorf("error scraping %s: %w", r.Request.URL.String(), err))
	})

	for _, url := range urls {
		if err := c.Visit(url); err != nil {
			saveError(fmt.Errorf("error queueing %s: %w", url, err))
		}
	}

	c.Wait()
	close(pages)
	writerWG.Wait()

	if writerErr != nil {
		return writerErr
	}

	errorLock.Lock()
	defer errorLock.Unlock()

	return scrapeErr
}

var citationRe = regexp.MustCompile(`\[\d+\]`)

// clean removes citation markers like [1] and collapses whitespace
func clean(content string) string {
	content = citationRe.ReplaceAllString(content, "")
	return strings.Join(strings.Fields(content), " ")
}
