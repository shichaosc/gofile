package main

import (
	"./postgresql"
	"./rabbitmq"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"strings"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

var p = rabbitmq.Producer{
	Exchange: "colly",
	Qname:    "qidian.book.content",
}

func main() {
	rabbitmq.Receive("colly.queue.index", collyQiDianIndex)
}

func collyQiDianIndex(url string) {

	c := colly.NewCollector()
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})
	var b = new(postgresql.Book)
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visited", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

	c.OnHTML("a[class='writer']", func(e *colly.HTMLElement) {
		//获得作者
		fmt.Println(e.Text)
	})

	c.OnHTML("div h1 em", func(e *colly.HTMLElement) {
		//获得书名
		fmt.Println(e.Text)
	})

	c.OnHTML("div[class='volume-wrap'] div[class='volume'] ul li a", func(e *colly.HTMLElement) {
		//获得章节名, url
		fmt.Println(e.Text)
		book_url := e.Attr("href")
		if strings.Index(book_url, "https:") == -1 {
			book_url = "https:" + book_url
		}
		p.Publish(book_url)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	c.Visit(url)
	fmt.Println(b)
}
