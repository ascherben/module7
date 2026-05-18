package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	output := "output.jl"

	start := time.Now()

	err := Scrape(output)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Scraping complete. Results saved to", output)
	fmt.Println("Runtime:", time.Since(start))
}
