package main

import (
	"./rabbitmq"
	"fmt"
	"github.com/gocolly/colly"
	"log"
)

//解析首页
func main() {
	var p = new(rabbitmq.Producer)
	p.Exchange = "colly"
	p.Qname = "colly.queue.index"

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visited", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

	c.OnHTML("div[class='book-list'] ul a[class='name']", func(e *colly.HTMLElement) {
		//三江
		book_url := e.Attr("href")
		p.Publish(book_url)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})
	c.Visit("https://www.qidian.com/")
}
