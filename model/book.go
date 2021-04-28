package model

import (
	"time"
)

// BookBook [...]
type BookBook struct {
	ID               int64      `gorm:"primaryKey;column:id;type:bigint;not null" json:"-"`
	BookName         string     `gorm:"column:book_name;type:varchar(64);not null" json:"bookName"`
	BookIntroduction string     `gorm:"column:book_introduction;type:text;not null" json:"bookIntroduction"`
	CoverImg         string     `gorm:"column:cover_img;type:varchar(256);not null" json:"coverImg"`
	TypeID           int        `gorm:"index:book_book_book_type_type_id;column:type_id;type:int" json:"typeId"`
	BookType         BookType   `gorm:"joinForeignKey:type_id;foreignKey:id" json:"bookTypeList"`
	AutherID         int        `gorm:"index:book_book_user_auther_auther_id;column:auther_id;type:int;not null" json:"autherId"`
	UserAuther       UserAuther `gorm:"joinForeignKey:auther_id;foreignKey:id" json:"userAutherList"`
	OnDelete         int        `gorm:"column:on_delete;type:int;default:0" json:"onDelete"`
	CreatedOn        time.Time  `gorm:"column:created_on;type:timestamp;default:CURRENT_TIMESTAMP" json:"createdOn"`
	UpdatedOn        time.Time  `gorm:"column:updated_on;type:timestamp;default:CURRENT_TIMESTAMP" json:"updatedOn"`
}

// TableName get sql table name.获取数据库表名
func (m *BookBook) TableName() string {
	return "book_book"
}

// BookBookColumns get sql column name.获取数据库列名
var BookBookColumns = struct {
	ID               string
	BookName         string
	BookIntroduction string
	CoverImg         string
	TypeID           string
	AutherID         string
	OnDelete         string
	CreatedOn        string
	UpdatedOn        string
}{
	ID:               "id",
	BookName:         "book_name",
	BookIntroduction: "book_introduction",
	CoverImg:         "cover_img",
	TypeID:           "type_id",
	AutherID:         "auther_id",
	OnDelete:         "on_delete",
	CreatedOn:        "created_on",
	UpdatedOn:        "updated_on",
}

// BookContent [...]
type BookContent struct {
	ID        int64     `gorm:"primaryKey;column:id;type:bigint;not null" json:"-"`
	Content   string    `gorm:"column:content;type:text;not null" json:"content"`
	BookID    int64     `gorm:"column:book_id;type:bigint;not null" json:"bookId"`
	TitleID   int64     `gorm:"column:title_id;type:bigint;not null" json:"titleId"`
	CreatedOn time.Time `gorm:"column:created_on;type:timestamp;default:CURRENT_TIMESTAMP" json:"createdOn"`
	UpdatedOn time.Time `gorm:"column:updated_on;type:timestamp;default:CURRENT_TIMESTAMP" json:"updatedOn"`
}

// TableName get sql table name.获取数据库表名
func (m *BookContent) TableName() string {
	return "book_content"
}

// BookContentColumns get sql column name.获取数据库列名
var BookContentColumns = struct {
	ID        string
	Content   string
	TitleID   string
	CreatedOn string
	UpdatedOn string
}{
	ID:        "id",
	Content:   "content",
	TitleID:   "title_id",
	CreatedOn: "created_on",
	UpdatedOn: "updated_on",
}

// BookReptileSource [...]
type BookReptileSource struct {
	ID        int64     `gorm:"primaryKey;column:id;type:bigint;not null" json:"-"`
	ContentID int64     `gorm:"column:content_id;type:bigint" json:"contentId"`
	Source    string    `gorm:"column:source;type:varchar(128)" json:"source"`
	SourceURL string    `gorm:"column:source_url;type:varchar(256)" json:"sourceUrl"` // url
	CreatedOn time.Time `gorm:"column:created_on;type:timestamp;default:CURRENT_TIMESTAMP" json:"createdOn"`
	UpdatedOn time.Time `gorm:"column:updated_on;type:timestamp;default:CURRENT_TIMESTAMP" json:"updatedOn"`
}

// TableName get sql table name.获取数据库表名
func (m *BookReptileSource) TableName() string {
	return "book_reptile_source"
}

// BookReptileSourceColumns get sql column name.获取数据库列名
var BookReptileSourceColumns = struct {
	ID        string
	ContentID string
	Source    string
	SourceURL string
	CreatedOn string
	UpdatedOn string
}{
	ID:        "id",
	ContentID: "content_id",
	Source:    "source",
	SourceURL: "source_url",
	CreatedOn: "created_on",
	UpdatedOn: "updated_on",
}

// BookTitle [...]
type BookTitle struct {
	ID        int64     `gorm:"primaryKey;column:id;type:bigint;not null" json:"-"`
	TitleName string    `gorm:"column:title_name;type:varchar(64);not null" json:"titleName"`
	Sort      int       `gorm:"column:sort;type:int;not null" json:"sort"`
	BookID    int64     `gorm:"index:book_book_title_book_book_id;column:book_id;type:bigint;not null" json:"bookId"`
	BookBook  BookBook  `gorm:"joinForeignKey:book_id;foreignKey:id" json:"bookBookList"`
	CreatedOn time.Time `gorm:"column:created_on;type:timestamp;default:CURRENT_TIMESTAMP" json:"createdOn"`
	UpdatedOn time.Time `gorm:"column:updated_on;type:timestamp;default:CURRENT_TIMESTAMP" json:"updatedOn"`
}

// TableName get sql table name.获取数据库表名
func (m *BookTitle) TableName() string {
	return "book_title"
}

// BookTitleColumns get sql column name.获取数据库列名
var BookTitleColumns = struct {
	ID        string
	TitleName string
	Sort      string
	BookID    string
	CreatedOn string
	UpdatedOn string
}{
	ID:        "id",
	TitleName: "title_name",
	Sort:      "sort",
	BookID:    "book_id",
	CreatedOn: "created_on",
	UpdatedOn: "updated_on",
}

// BookType [...]
type BookType struct {
	ID        int       `gorm:"primaryKey;column:id;type:int;not null" json:"-"`
	TypeName  string    `gorm:"column:type_name;type:varchar(32);not null" json:"typeName"`
	Sort      int       `gorm:"column:sort;type:int;not null" json:"sort"`
	CreatedOn time.Time `gorm:"column:created_on;type:timestamp;default:CURRENT_TIMESTAMP" json:"createdOn"`
	UpdatedOn time.Time `gorm:"column:updated_on;type:timestamp;default:CURRENT_TIMESTAMP" json:"updatedOn"`
}

// TableName get sql table name.获取数据库表名
func (m *BookType) TableName() string {
	return "book_type"
}

// BookTypeColumns get sql column name.获取数据库列名
var BookTypeColumns = struct {
	ID        string
	TypeName  string
	Sort      string
	CreatedOn string
	UpdatedOn string
}{
	ID:        "id",
	TypeName:  "type_name",
	Sort:      "sort",
	CreatedOn: "created_on",
	UpdatedOn: "updated_on",
}

// UserAuther [...]
type UserAuther struct {
	ID           int       `gorm:"primaryKey;column:id;type:int;not null" json:"-"` // id
	Username     string    `gorm:"column:username;type:varchar(64);not null" json:"username"`
	Realname     string    `gorm:"column:realname;type:varchar(32)" json:"realname"`
	Introduction string    `gorm:"column:introduction;type:text" json:"introduction"`
	Avatar       string    `gorm:"column:avatar;type:varchar(256)" json:"avatar"`
	Gender       int       `gorm:"index:gender;column:gender;type:int" json:"gender"`
	OnDelete     int       `gorm:"column:on_delete;type:int" json:"onDelete"`
	CreatedOn    time.Time `gorm:"column:created_on;type:timestamp;default:CURRENT_TIMESTAMP" json:"createdOn"`
	UpdatedOn    time.Time `gorm:"column:updated_on;type:timestamp;default:CURRENT_TIMESTAMP" json:"updatedOn"`
}

// TableName get sql table name.获取数据库表名
func (m *UserAuther) TableName() string {
	return "user_auther"
}

// UserAutherColumns get sql column name.获取数据库列名
var UserAutherColumns = struct {
	ID           string
	Username     string
	Realname     string
	Introduction string
	Avatar       string
	Gender       string
	OnDelete     string
	CreatedOn    string
	UpdatedOn    string
}{
	ID:           "id",
	Username:     "username",
	Realname:     "realname",
	Introduction: "introduction",
	Avatar:       "avatar",
	Gender:       "gender",
	OnDelete:     "on_delete",
	CreatedOn:    "created_on",
	UpdatedOn:    "updated_on",
}
