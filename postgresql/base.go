package postgresql

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Book struct {
	Title   string
	Content string
	Order   int
}

const (
	HOST     = "139.9.92.116"
	PORT     = 5432
	USER     = "postgres"
	PASSWORD = "Shichao123"
	DBNAME   = "postgres"
)

//Db数据库连接池
var DB *sql.DB

//注意方法名大写，就是public
func InitDB() {
	//path := strings.Join([]string{USERNAME, ":", PASSWORD, "@tcp(", SERVER, ":", PORT, ")/", DATABASE, "?charset=utf8"}, "")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		HOST, PORT, USER, PASSWORD, DBNAME)

	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("postgres", psqlInfo)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(20)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("opon database fail")
		return
	}
	fmt.Println("connnect success")
}

func InsertBook(book Book) bool {
	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx fail")
		return false
	}
	//准备sql语句
	stmt, err := tx.Prepare("INSERT INTO book_content (`title`, `content`, `order`, `book_id`) VALUES (?, ?, ?, ?)")
	if err != nil {
		fmt.Println("Prepare fail")
		return false
	}
	//将参数传递到sql语句中并且执行
	res, err := stmt.Exec(book.Title, book.Content, book.Order, 1)
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

func init() {
	InitDB()
}
