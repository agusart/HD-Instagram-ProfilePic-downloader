package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("invalid argument")
	}

	username := os.Args[1]
	res, err := http.Get(
		fmt.Sprintf("https://www.instadp.com/fullsize/%s", username))
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".download-btn").Each(func(i int, s *goquery.Selection) {
		href, ok := s.Attr("href")

		if ok {
			fmt.Printf(href)
		}
	})

}
