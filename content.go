package main

import (
	"./postgresql"
	"./rabbitmq"
	"fmt"
	"github.com/gocolly/colly"
	"log"
)

var order = 0

func main() {
	rabbitmq.Receive("qidian.book.content", collyContent)
}

//解析内容
func collyContent(url string) {
	order++
	var book = postgresql.BookContent{}
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visited", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("finished", r.Request.URL)

	})

	c.OnHTML("div[class='main-text-wrap'] div[class='info f1'] a", func(e *colly.HTMLElement) {
		//书名跟作者
		fmt.Println(e.Text)
	})

	c.OnHTML("h3[class='j_chapterName']", func(e *colly.HTMLElement) {
		//章节名
		chapter := e.Text
		book.Title = chapter
		fmt.Println(chapter)
	})

	c.OnHTML("div[class='read-content j_readContent']", func(e *colly.HTMLElement) {
		//内容
		content := e.Text
		book.Content = content
		fmt.Println(content)
	})

	c.OnScraped(func(r *colly.Response) {
		postgresql.InsertBook(book)
		fmt.Println("Finished", r.Request.URL)
	})
	c.Visit(url)

}
