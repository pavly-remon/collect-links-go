package main

import (
	"fmt"
	"strings"

	s "com.picode/collect-links-go/ScrapLinks"
)

func main() {
	for {
		url, err := GetUserInput("Enter URL You want to Scan: ")
		if err != nil {
			fmt.Println("Enter a Valid URL")
			continue
		}
		fmt.Printf("Start Processing for URL: %s...\n", url)
		links, err := s.ScrapLinks(url)
		if err != nil {
			fmt.Println(err)
			continue
		}
		splittedUrl := strings.Split(url, ".")
		WriteFile(splittedUrl[1], links)
		fmt.Printf("JOB DONE for URL: %s...\n", url)

	}
}
