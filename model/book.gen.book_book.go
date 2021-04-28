package model

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _BookBookMgr struct {
	*_BaseMgr
}

// BookBookMgr open func
func BookBookMgr(db *gorm.DB) *_BookBookMgr {
	if db == nil {
		panic(fmt.Errorf("BookBookMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_BookBookMgr{_BaseMgr: &_BaseMgr{DB: db.Table("book_book"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_BookBookMgr) GetTableName() string {
	return "book_book"
}

// Get 获取
func (obj *_BookBookMgr) Get() (result BookBook, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("book_type").Where("id = ?", result.TypeID).Find(&result.BookType).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("user_auther").Where("id = ?", result.AutherID).Find(&result.UserAuther).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_BookBookMgr) Gets() (results []*BookBook, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("book_type").Where("id = ?", results[i].TypeID).Find(&results[i].BookType).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_auther").Where("id = ?", results[i].AutherID).Find(&results[i].UserAuther).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_BookBookMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithBookName book_name获取
func (obj *_BookBookMgr) WithBookName(bookName string) Option {
	return optionFunc(func(o *options) { o.query["book_name"] = bookName })
}

// WithBookIntroduction book_introduction获取
func (obj *_BookBookMgr) WithBookIntroduction(bookIntroduction string) Option {
	return optionFunc(func(o *options) { o.query["book_introduction"] = bookIntroduction })
}

// WithCoverImg cover_img获取
func (obj *_BookBookMgr) WithCoverImg(coverImg string) Option {
	return optionFunc(func(o *options) { o.query["cover_img"] = coverImg })
}

// WithTypeID type_id获取
func (obj *_BookBookMgr) WithTypeID(typeID int) Option {
	return optionFunc(func(o *options) { o.query["type_id"] = typeID })
}

// WithAutherID auther_id获取
func (obj *_BookBookMgr) WithAutherID(autherID int) Option {
	return optionFunc(func(o *options) { o.query["auther_id"] = autherID })
}

// WithOnDelete on_delete获取
func (obj *_BookBookMgr) WithOnDelete(onDelete int) Option {
	return optionFunc(func(o *options) { o.query["on_delete"] = onDelete })
}

// WithCreatedOn created_on获取
func (obj *_BookBookMgr) WithCreatedOn(createdOn time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_on"] = createdOn })
}

// WithUpdatedOn updated_on获取
func (obj *_BookBookMgr) WithUpdatedOn(updatedOn time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_on"] = updatedOn })
}

// GetByOption 功能选项模式获取
func (obj *_BookBookMgr) GetByOption(opts ...Option) (result BookBook, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.Table(obj.GetTableName()).Where(options.query).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("book_type").Where("id = ?", result.TypeID).Find(&result.BookType).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("user_auther").Where("id = ?", result.AutherID).Find(&result.UserAuther).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_BookBookMgr) GetByOptions(opts ...Option) (results []*BookBook, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.Table(obj.GetTableName()).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("book_type").Where("id = ?", results[i].TypeID).Find(&results[i].BookType).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_auther").Where("id = ?", results[i].AutherID).Find(&results[i].UserAuther).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_BookBookMgr) GetFromID(id int64) (result BookBook, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("book_type").Where("id = ?", result.TypeID).Find(&result.BookType).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("user_auther").Where("id = ?", result.AutherID).Find(&result.UserAuther).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromID 批量查找
func (obj *_BookBookMgr) GetBatchFromID(ids []int64) (results []*BookBook, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("book_type").Where("id = ?", results[i].TypeID).Find(&results[i].BookType).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_auther").Where("id = ?", results[i].AutherID).Find(&results[i].UserAuther).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromBookName 通过book_name获取内容
func (obj *_BookBookMgr) GetFromBookName(bookName string) (result BookBook, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("book_name = ?", bookName).First(&result).Error
	return
}

// GetBatchFromBookName 批量查找
func (obj *_BookBookMgr) GetBatchFromBookName(bookNames []string) (results []*BookBook, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("book_name IN (?)", bookNames).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("book_type").Where("id = ?", results[i].TypeID).Find(&results[i].BookType).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_auther").Where("id = ?", results[i].AutherID).Find(&results[i].UserAuther).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromBookIntroduction 通过book_introduction获取内容
func (obj *_BookBookMgr) GetFromBookIntroduction(bookIntroduction string) (results []*BookBook, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("book_introduction = ?", bookIntroduction).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("book_type").Where("id = ?", results[i].TypeID).Find(&results[i].BookType).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_auther").Where("id = ?", results[i].AutherID).Find(&results[i].UserAuther).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromBookIntroduction 批量查找
func (obj *_BookBookMgr) GetBatchFromBookIntroduction(bookIntroductions []string) (results []*BookBook, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("book_introduction IN (?)", bookIntroductions).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("book_type").Where("id = ?", results[i].TypeID).Find(&results[i].BookType).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_auther").Where("id = ?", results[i].AutherID).Find(&results[i].UserAuther).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCoverImg 通过cover_img获取内容
func (obj *_BookBookMgr) GetFromCoverImg(coverImg string) (results []*BookBook, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("cover_img = ?", coverImg).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("book_type").Where("id = ?", results[i].TypeID).Find(&results[i].BookType).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_auther").Where("id = ?", results[i].AutherID).Find(&results[i].UserAuther).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCoverImg 批量查找
func (obj *_BookBookMgr) GetBatchFromCoverImg(coverImgs []string) (results []*BookBook, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("cover_img IN (?)", coverImgs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("book_type").Where("id = ?", results[i].TypeID).Find(&results[i].BookType).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_auther").Where("id = ?", results[i].AutherID).Find(&results[i].UserAuther).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromTypeID 通过type_id获取内容
func (obj *_BookBookMgr) GetFromTypeID(typeID int) (results []*BookBook, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("type_id = ?", typeID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("book_type").Where("id = ?", results[i].TypeID).Find(&results[i].BookType).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_auther").Where("id = ?", results[i].AutherID).Find(&results[i].UserAuther).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromTypeID 批量查找
func (obj *_BookBookMgr) GetBatchFromTypeID(typeIDs []int) (results []*BookBook, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("type_id IN (?)", typeIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("book_type").Where("id = ?", results[i].TypeID).Find(&results[i].BookType).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_auther").Where("id = ?", results[i].AutherID).Find(&results[i].UserAuther).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromAutherID 通过auther_id获取内容
func (obj *_BookBookMgr) GetFromAutherID(autherID int) (results []*BookBook, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("auther_id = ?", autherID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("book_type").Where("id = ?", results[i].TypeID).Find(&results[i].BookType).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_auther").Where("id = ?", results[i].AutherID).Find(&results[i].UserAuther).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromAutherID 批量查找
func (obj *_BookBookMgr) GetBatchFromAutherID(autherIDs []int) (results []*BookBook, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("auther_id IN (?)", autherIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("book_type").Where("id = ?", results[i].TypeID).Find(&results[i].BookType).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_auther").Where("id = ?", results[i].AutherID).Find(&results[i].UserAuther).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromOnDelete 通过on_delete获取内容
func (obj *_BookBookMgr) GetFromOnDelete(onDelete int) (results []*BookBook, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("on_delete = ?", onDelete).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("book_type").Where("id = ?", results[i].TypeID).Find(&results[i].BookType).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_auther").Where("id = ?", results[i].AutherID).Find(&results[i].UserAuther).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromOnDelete 批量查找
func (obj *_BookBookMgr) GetBatchFromOnDelete(onDeletes []int) (results []*BookBook, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("on_delete IN (?)", onDeletes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("book_type").Where("id = ?", results[i].TypeID).Find(&results[i].BookType).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_auther").Where("id = ?", results[i].AutherID).Find(&results[i].UserAuther).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreatedOn 通过created_on获取内容
func (obj *_BookBookMgr) GetFromCreatedOn(createdOn time.Time) (results []*BookBook, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("created_on = ?", createdOn).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("book_type").Where("id = ?", results[i].TypeID).Find(&results[i].BookType).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_auther").Where("id = ?", results[i].AutherID).Find(&results[i].UserAuther).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreatedOn 批量查找
func (obj *_BookBookMgr) GetBatchFromCreatedOn(createdOns []time.Time) (results []*BookBook, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("created_on IN (?)", createdOns).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("book_type").Where("id = ?", results[i].TypeID).Find(&results[i].BookType).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_auther").Where("id = ?", results[i].AutherID).Find(&results[i].UserAuther).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdatedOn 通过updated_on获取内容
func (obj *_BookBookMgr) GetFromUpdatedOn(updatedOn time.Time) (results []*BookBook, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("updated_on = ?", updatedOn).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("book_type").Where("id = ?", results[i].TypeID).Find(&results[i].BookType).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_auther").Where("id = ?", results[i].AutherID).Find(&results[i].UserAuther).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdatedOn 批量查找
func (obj *_BookBookMgr) GetBatchFromUpdatedOn(updatedOns []time.Time) (results []*BookBook, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("updated_on IN (?)", updatedOns).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("book_type").Where("id = ?", results[i].TypeID).Find(&results[i].BookType).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_auther").Where("id = ?", results[i].AutherID).Find(&results[i].UserAuther).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_BookBookMgr) FetchByPrimaryKey(id int64) (result BookBook, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("book_type").Where("id = ?", result.TypeID).Find(&result.BookType).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.New().Table("user_auther").Where("id = ?", result.AutherID).Find(&result.UserAuther).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByBookBookBookTypeTypeID  获取多个内容
func (obj *_BookBookMgr) FetchIndexByBookBookBookTypeTypeID(typeID int) (results []*BookBook, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("type_id = ?", typeID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("book_type").Where("id = ?", results[i].TypeID).Find(&results[i].BookType).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_auther").Where("id = ?", results[i].AutherID).Find(&results[i].UserAuther).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByBookBookUserAutherAutherID  获取多个内容
func (obj *_BookBookMgr) FetchIndexByBookBookUserAutherAutherID(autherID int) (results []*BookBook, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("auther_id = ?", autherID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("book_type").Where("id = ?", results[i].TypeID).Find(&results[i].BookType).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.New().Table("user_auther").Where("id = ?", results[i].AutherID).Find(&results[i].UserAuther).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
