package main

import (
	"encoding/csv"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"os"
)

type Data struct {
	Title string `json:"title"`
	Link string `json:link`
}

func main() {
	fileName := "viblo_trending.csv"
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Could not create %s", fileName)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Title", "Link"})

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnHTML(".post-feed .link", func(e *colly.HTMLElement) {
		data := Data{}
		data.Title = e.Text
		data.Link = "https://viblo.asia" +e.Attr("href")
		if data.Title == "" || data.Link == "https://viblo.asia" {
			return
		}
		//fmt.Printf("Title: %s \nLink: %s \n", res.Title, res.Link)
		writer.Write([]string{data.Title, data.Link})
	})

	for i := 1; i < 5; i++ {
		fullURL := fmt.Sprintf("https://viblo.asia/trending?page=%d", i)
		c.Visit(fullURL)
	}
	c.Wait()
}