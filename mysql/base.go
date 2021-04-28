package mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"reptile_book/config"
)

//
//type BookContent struct {
//	Id        big.Int
//	TitleId   big.Int
//	Content   string
//	CreatedOn time.Time
//	UpdatedOn time.Time
//}
//
//type BookTitle struct {
//	Id        big.Int
//	TitleName string
//	Sort      int
//	BookId    big.Int
//	CreatedOn time.Time
//	UpdatedOn time.Time
//}
//
//type Book struct {
//	Id        big.Int
//	BookName  string
//	TypeId    int
//	CreatedOn time.Time
//	UpdatedOn time.Time
//}
//
//type BookType struct {
//	Id        int
//	TypeName  string
//	CreatedOn time.Time
//	UpdatedOn time.Time
//}

const (
	USERNAME = "book"
	PASSWORD = "bookPWD"
	SERVER   = "82.156.237.178"
	PORT     = "3306"
	DATABASE = "book"
)

//Db数据库连接池
//var DB *sql.DB
//
////注意方法名大写，就是public
//func InitDB() {
//	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
//	path := strings.Join([]string{USERNAME, ":", PASSWORD, "@tcp(", SERVER, ":", PORT, ")/", DATABASE, "?charset=utf8"}, "")
//
//	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
//	DB, _ = sql.Open("mysql", path)
//	//设置数据库最大连接数
//	DB.SetConnMaxLifetime(20)
//	//设置上数据库最大闲置连接数
//	DB.SetMaxIdleConns(10)
//	//验证连接
//	if err := DB.Ping(); err != nil {
//		fmt.Println("opon database fail")
//		return
//	}
//	fmt.Println("connnect success")
//}

//func InsertBook(book Book) bool {
//	//开启事务
//	tx, err := DB.Begin()
//	if err != nil {
//		fmt.Println("tx fail")
//		return false
//	}
//	//准备sql语句
//	stmt, err := tx.Prepare("INSERT INTO book_content (`title_id`, `content`, `sort`, `book_id`) VALUES (?, ?, ?, ?)")
//	if err != nil {
//		fmt.Println("Prepare fail")
//		return false
//	}
//	//将参数传递到sql语句中并且执行
//	res, err := stmt.Exec(book.Title, book.Content, book.Order, 1)
//	if err != nil {
//		fmt.Println("Exec fail")
//		return false
//	}
//	//将事务提交
//	tx.Commit()
//	//获得上一个插入自增的id
//	fmt.Println(res.LastInsertId())
//	return true
//}

//func init() {
//	InitDB()
//}

var DB *gorm.DB

func initDB() {
	//DB, err := gorm.Open(config.DBConfig.Dialect, config.DBConfig.URL)
	DB, _ = gorm.Open("mysql", "book:bookPWD@tcp(82.156.237.178:3306)/book?charset=utf8&parseTime=True")
	//defer database.Close()

	DB.DB().SetMaxIdleConns(config.DBConfig.MaxIdleConns)
	DB.DB().SetMaxOpenConns(config.DBConfig.MaxOpenConns)
	DB.SingularTable(true)
}

func init() {
	initDB()
}
