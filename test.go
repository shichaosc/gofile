package main

import (
	"fmt"
	"reptile_book/model"
	"reptile_book/mysql"
	"encoding/json"
)

func main() {
	sort := mysql.QueryBookTitleSort(1)
	fmt.Println("success")
	fmt.Println(sort)
	book_type, err := model.BookTypeMgr(mysql.DB).GetFromID(2)
	if err != nil {
		println(err.Error())
	}
	fmt.Print(book_type)
	jsonBytes, _ := json.Marshal(book_type)
	fmt.Println(string(jsonBytes))
}
