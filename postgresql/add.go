package postgresql

import (
	"fmt"
	"log"
	"reptile_book/model"
	_ "reptile_book/model"
)

func InsertBookType(bookType model.BookType) bool {
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx fail")
		return false
	}
	stmt, err := tx.Prepare("INSERT INTO book_type (type_name, created_on, updated_on) VALUES ($1,$2,$3)")
	if err != nil {
		fmt.Println("Prepare fail")
		log.Fatalf("publish msg fail: err=%s", err)
		return false
	}
	res, err := stmt.Exec(bookType.TypeName, bookType.CreatedOn, bookType.UpdatedOn)
	log.Fatalf("result=%s", res)
	if err != nil {
		fmt.Println("Exec fail")
		return false
	}
	tx.Commit()
	fmt.Println(res.LastInsertId())
	return true
}

func InsertBook(book model.Book) bool {
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx fail")
		return false
	}
	stmt, err := tx.Prepare("INSERT INTO book_book (book_name, type_id, created_on, updated_on) VALUES ($1,$2,$3,$4)")
	if err != nil {
		fmt.Println("Prepare fail")
		log.Fatalf("publish msg fail: err=%s", err)
		return false
	}
	res, err := stmt.Exec(book.BookName, book.TypeId, book.CreatedOn, book.UpdatedOn)
	log.Fatalf("result=%s", res)
	if err != nil {
		fmt.Println("Exec fail")
		return false
	}
	tx.Commit()
	fmt.Println(res.LastInsertId())
	return true
}

func InsertBookTitle(bookTitle model.BookTitle) bool {
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx fail")
		return false
	}
	stmt, err := tx.Prepare("INSERT INTO book_title (title_name, book_id, created_on, updated_on) VALUES ($1,$2,$3, $4)")
	if err != nil {
		fmt.Println("Prepare fail")
		log.Fatalf("publish msg fail: err=%s", err)
		return false
	}
	res, err := stmt.Exec(bookTitle.TitleName, bookTitle.BookId, bookTitle.CreatedOn, bookTitle.UpdatedOn)
	log.Fatalf("result=%s", res)
	if err != nil {
		fmt.Println("Exec fail")
		return false
	}
	tx.Commit()
	fmt.Println(res.LastInsertId())
	return true
}

func InsertBookContent(bookContent model.BookContent) bool {
	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx fail")
		return false
	}
	//准备sql语句
	stmt, err := tx.Prepare("INSERT INTO book_content (title_id, content, created_on, updated_on) VALUES ($1,$2,$3,$4)")
	if err != nil {
		fmt.Println("Prepare fail")
		log.Fatalf("publish msg fail: err=%s", err)
		return false
	}
	//将参数传递到sql语句中并且执行
	res, err := stmt.Exec(bookContent.TitleId, bookContent.Content, bookContent.CreatedOn, bookContent.UpdatedOn)
	log.Fatalf("result=%s", res)
	if err != nil {
		fmt.Println("Exec fail")
		return false
	}
	//将事务提交
	tx.Commit()
	//获得上一个插入自增的id
	fmt.Println(res.LastInsertId())
	return true
}
