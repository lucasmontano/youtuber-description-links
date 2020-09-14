package main

import (
	"fmt"
	"os"
)

func main() {
	links := ExtractVideoLinks(os.Args[1])
	if len(links) == 0 {
		fmt.Println("No links found")
	}
	for _, link := range links {
		fmt.Println(link)
	}
}
