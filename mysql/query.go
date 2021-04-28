package mysql

import (
	_ "github.com/jinzhu/gorm"
	_ "reptile_book/model"
)

//func queryBook(bookName string) {
//	err := DB.QueryRow("SELECT * FROM book_book where book_name=$1", bookName)
//
//	if err != nil {
//		fmt.Println(err)
//	}
//	//fmt.Printf()
//	//defer rows.Close()
//
//	//for rows.Next() {
//	//	p := new(Book)
//	//	err := rows.Scan(&p.uid, &p.username, &p.departname,&p.Created)
//	//	if err != nil {
//	//		fmt.Println(err)
//	//	}
//	//	fmt.Println(p.uid, p.username, p.departname,p.Created)
//	//}
//}

//func QueryBookTitleSort(bookId int64) int64 {
//
//	res, _ := model.BookTitleMgr(DB).GetMaxSortFromBookID(bookId)
//	return res
//}
