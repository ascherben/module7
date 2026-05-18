# Go Web Scraper with Colly

## Overview

This project builds a Go-based web scraper using Colly. The scraper visits ten Wikipedia pages related to robotics, intelligent systems, and artificial intelligence. It extracts text from each page and saves the results in a JSON Lines file. The runtime is then compared to a Python/Scrapy example that scrapes the same webpages.

The output file is: `output.jl`

Each line in the file contains one JSON object for one scraped Wikipedia page. Each object includes the page URL, extracted text, and a timestamp.

The scraper visits the following Wikipedia pages:

```
https://en.wikipedia.org/wiki/Robotics
https://en.wikipedia.org/wiki/Robot
https://en.wikipedia.org/wiki/Reinforcement_learning
https://en.wikipedia.org/wiki/Robot_Operating_System
https://en.wikipedia.org/wiki/Intelligent_agent
https://en.wikipedia.org/wiki/Software_agent
https://en.wikipedia.org/wiki/Robotic_process_automation
https://en.wikipedia.org/wiki/Chatbot
https://en.wikipedia.org/wiki/Applications_of_artificial_intelligence
https://en.wikipedia.org/wiki/Android_(robot)
```

### Project Files

- .gitignore
- ai/
  - claude.txt
- go.mod
- go.sum
- main.go
- scraper.go
- scraper_test.go
- output.jl
- README.md

### Requirements

- Go 1.26 or higher

### Download the Repository

Download or clone the repository

### Run with Go

```bash
go run .
```

Results are written to `output.jl` 

### Build an Executable

#### Windows

```powershell
go build -o main.exe .
.\main.exe
```

#### macOS / Linux

```bash
go build -o main .
./main
```

### Testing

```bash
go test ./...
```

The tests check the text cleaning function and confirm that the Page function creates the expected JSON fields.

## Timing Comparison

I also compared the Go/Colly scraper with the provided Python/Scrapy example using the same ten Wikipedia URLs and averaged the runtimes. 

| Version | Runtime |
|---|---:|
| Python/Scrapy | ~ 16.25 seconds |
| Go/Colly | ~ 5.53 seconds |


The scraper successfully creates a JSON Lines file with one JSON object per Wikipedia page. The output includes the page URL, extracted text, and timestamp.


## Recommendation

The Go/Colly scraper is a good replacement for the Python/Scrapy example for this assignment. It uses the same URL list, produces JSON Lines output, and runs faster in this test. Go also provides a clear structure for concurrency using Colly's asynchronous collector.

## GenAI Tools and Sources

**Sources**

- https://github.com/seversky/gachifinder/blob/master/scrape/scrape.go

- https://github.com/gocolly/colly

**AI**

GenAI was used to help understand Colly syntax, debug code, create simple unit tests, and edit the README. The final code was reviewed and edited before submission.


