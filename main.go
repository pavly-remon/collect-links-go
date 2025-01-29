package main

import (
	"fmt"
	"strings"

	"com.picode/collect-links-go/link-collector"
)

func main() {
	for {
		inputUrl, err := GetUserInput("Enter URL You want to Scan: ")

		if err != nil {
			fmt.Println("Enter a Valid URL")
			continue
		}

		fmt.Printf("Start Processing for URL: %s...\n", inputUrl)

		linkCollector := link_collector.New(inputUrl)

		err = linkCollector.CollectLinks()
		if err != nil {
			fmt.Println(err)
			continue
		}
		splittedUrl := strings.Split(inputUrl, ".")
		WriteFile(splittedUrl[1], linkCollector.GetLinks())
		fmt.Printf("JOB DONE for URL: %s...\n", inputUrl)

	}
}
