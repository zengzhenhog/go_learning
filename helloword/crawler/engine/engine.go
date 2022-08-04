package engine

import (
	"helloword/crawler/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request
	requests = append(requests, seeds...)

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		log.Printf("fetching %s", r.URL)
		body, err := fetcher.Fetch(r.URL)
		if err != nil {
			log.Printf("Fetcher: error fetching URL: %s: %v", r.URL, err)
		}

		parserResult := r.ParserFunc(body)
		requests = append(requests, parserResult.Requests...)

		for _, item := range parserResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}
