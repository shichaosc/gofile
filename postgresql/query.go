package postgresql

import (
	"fmt"
)

func queryBook(bookName string) {
	err := DB.QueryRow("SELECT * FROM book_book where book_name=$1", bookName)

	if err != nil {
		fmt.Println(err)
	}
	//fmt.Printf()
	//defer rows.Close()

	//for rows.Next() {
	//	p := new(Book)
	//	err := rows.Scan(&p.uid, &p.username, &p.departname,&p.Created)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	fmt.Println(p.uid, p.username, p.departname,p.Created)
	//}
}
