package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	_ "strings"
	_ "unicode/utf8"
	//"github.com/shen100/golang123/utils"
)

var jsonData map[string]interface{}

func initJSON() {
	bytes, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
		os.Exit(-1)
	}

	configStr := string(bytes[:])
	reg := regexp.MustCompile(`/\*.*\*/`)

	configStr = reg.ReplaceAllString(configStr, "")
	bytes = []byte(configStr)

	if err := json.Unmarshal(bytes, &jsonData); err != nil {
		fmt.Println("invalid config: ", err.Error())
		os.Exit(-1)
	}
}

type dBConfig struct {
	Dialect         string
	Database        string
	User            string
	Password        string
	Host            string
	Port            int
	Charset         string
	URL             string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime int64
	Sslmode         string
}

// DBConfig 数据库相关配置
var DBConfig dBConfig

func initDB() {
	//utils.SetStructByJSON(&DBConfig, jsonData["database"].(map[string]interface{}))

	DBConfig.User = "book"
	DBConfig.Password = "bookPWD"
	DBConfig.Host = "82.156.237.178"
	DBConfig.Port = 3369
	DBConfig.Database = "book"
	DBConfig.Charset = "utf-8"

	//mysql数据库的连接方式
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		DBConfig.User, DBConfig.Password, DBConfig.Host, DBConfig.Port, DBConfig.Database, DBConfig.Charset)

	/**
	62         更改mysql数据库为postgresql
	63         具体连接方式为>
	64             host=myhost port=myport user=gorm dbname=gorm password=mypassword
	65      */
	//url := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
	//		 DBConfig.Host,
	//		DBConfig.Port,
	//		 DBConfig.User,
	//		 DBConfig.Database,
	//		 DBConfig.Password,
	//		 DBConfig.Sslmode)
	DBConfig.URL = url
}

type serverConfig struct {
	APIPoweredBy       string
	SiteName           string
	Host               string
	ImgHost            string
	Env                string
	LogDir             string
	LogFile            string
	APIPrefix          string
	UploadImgDir       string
	ImgPath            string
	MaxMultipartMemory int
	Port               int
	StatsEnabled       bool
	TokenSecret        string
	TokenMaxAge        int
	PassSalt           string
	LuosimaoVerifyURL  string
	LuosimaoAPIKey     string
	CrawlerName        string
	MailUser           string //域名邮箱账号
	MailPass           string //域名邮箱密码
	MailHost           string //smtp邮箱域名
	MailPort           int    //smtp邮箱端口
	MailFrom           string //邮件来源
	Github             string
	BaiduPushLink      string
}

func init() {
	//initJSON()
	initDB()
}
