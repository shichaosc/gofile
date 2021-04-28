package postgresql

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

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

func init() {
	InitDB()
}
