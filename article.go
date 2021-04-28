package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/bitly/go-simplejson"
	"github.com/gocolly/colly"
	"log"
	"reptile_book/model"
	"reptile_book/mysql"
	"reptile_book/rabbitmq"
	"strings"
)

var p = rabbitmq.Producer{
	Exchange: "colly",
	Qname:    "qidian.book.content",
}

func main() {
	rabbitmq.Receive("colly.queue.index", collyQiDianIndex)
}

type TitleInfo struct {
	TitleName string
	ContentUrl string
	Sort int
}

type BookInfo struct {
	AutherName string
	BookName string
	BookIntroduction string
	BookChapterNumber  int64
	BookWordNumber int64
	TitleInfo []TitleInfo
	TypeName string
}

type PublishTitle struct {
	ContentUrl string  `json:"contentUrl,string"`
	BookId   int64  `json:"bookId,int64"`
	TitleId  int64  `json:"titleId,int64"`
}

func collyQiDianIndex(url string) {

	fmt.Println(url)
	var sort = 0

	var bookInfo = BookInfo{}

	c := colly.NewCollector()
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visited", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

	c.OnHTML("a[class='writer']", func(e *colly.HTMLElement) {
		//获得作者
		fmt.Println(e.Text)
		bookInfo.AutherName = e.Text
	})

	c.OnHTML("div h1 em", func(e *colly.HTMLElement) {
		//获得书名
		fmt.Println(e.Text)
		bookInfo.BookName = e.Text
	})

	c.OnHTML("div[class='left-wrap fl'] div[class='book-info-detail'] div[class='book-intro'] p", func(e *colly.HTMLElement) {
		//获得书籍简介
		fmt.Println(e.Text)
		bookInfo.BookIntroduction = e.Text
	})

	c.OnHTML("div[class='volume-wrap'] div[class='volume'] ul[class='cf'] li a", func(e *colly.HTMLElement) {
		//获得章节名, url
		fmt.Println(e.Text)
		book_url := e.Attr("href")
		if strings.Index(book_url, "https:") == -1 {
			book_url = "https:" + book_url
		}
		titleInfo := TitleInfo{}
		titleInfo.TitleName = e.Text
		titleInfo.Sort = sort + 1
		titleInfo.ContentUrl = book_url
		bookInfo.TitleInfo = append(bookInfo.TitleInfo, titleInfo)
	})

	c.OnHTML("div[class='crumbs-nav center990  top-op'] span a:nth-child(6)", func(e *colly.HTMLElement) {
		fmt.Println("no error", e.Text)
		bookInfo.TypeName = e.Text
	})


	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
		auther, err := model.UserAutherMgr(mysql.DB).GetFromUsername(bookInfo.AutherName)
		if err != nil {
			auther = model.UserAuther{}
			auther.Username = bookInfo.AutherName
			model.UserAutherMgr(mysql.DB).Create(&auther)
		}

		bookType, err := model.BookTypeMgr(mysql.DB).GetFromTypeName(bookInfo.TypeName)

		if err != nil {
			bookType = model.BookType{}
			bookType.TypeName = bookInfo.TypeName
			model.BookTypeMgr(mysql.DB).Create(&bookType)

		}
		fmt.Println("autherId", auther.ID)
		book, bookErr := model.BookBookMgr(mysql.DB).GetFromBookName(bookInfo.BookName)

		if bookErr != nil {
			book = model.BookBook{}
			book.BookName = bookInfo.BookName
			book.AutherID = auther.ID
			book.BookIntroduction = bookInfo.BookIntroduction
			book.TypeID = bookType.ID
			model.BookBookMgr(mysql.DB).Create(&book)
		}

		for i := 0; i < len(bookInfo.TitleInfo); i++ {
			title, err := model.BookTitleMgr(mysql.DB).GetFromBookIDAndTitleName(book.ID, bookInfo.TitleInfo[i].TitleName)
			if err != nil {
				title = model.BookTitle{}
				title.TitleName = bookInfo.TitleInfo[i].TitleName
				title.BookID = book.ID
				maxTitle, err := model.BookTitleMgr(mysql.DB).GetMaxSortFromBookID(book.ID)
				maxSort := 0
				if err == nil {
					maxSort = maxTitle.Sort + 1
				}else{
					fmt.Println(err)
				}
				title.Sort = int(maxSort)
				model.BookTitleMgr(mysql.DB).Create(&title)
				fmt.Println("title id: ", title.ID)
			}

			var publishData = PublishTitle{}
			publishData.BookId = book.ID
			publishData.TitleId = title.ID
			publishData.ContentUrl = bookInfo.TitleInfo[i].ContentUrl
			data, _ := json.Marshal(&p)
			fmt.Println(string(data))
			p.Publish(data)
		}

	})

	c.Visit(url)
}
