package model

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _BookTitleMgr struct {
	*_BaseMgr
}

// BookTitleMgr open func
func BookTitleMgr(db *gorm.DB) *_BookTitleMgr {
	if db == nil {
		panic(fmt.Errorf("BookTitleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_BookTitleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("book_title"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_BookTitleMgr) GetTableName() string {
	return "book_title"
}

// Get 获取
func (obj *_BookTitleMgr) Get() (result BookTitle, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("book_book").Where("id = ?", result.BookID).Find(&result.BookBook).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_BookTitleMgr) Gets() (results []*BookTitle, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("book_book").Where("id = ?", results[i].BookID).Find(&results[i].BookBook).Error; err != nil { //
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
func (obj *_BookTitleMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithTitleName title_name获取
func (obj *_BookTitleMgr) WithTitleName(titleName string) Option {
	return optionFunc(func(o *options) { o.query["title_name"] = titleName })
}

// WithSort sort获取
func (obj *_BookTitleMgr) WithSort(sort int) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// WithBookID book_id获取
func (obj *_BookTitleMgr) WithBookID(bookID int64) Option {
	return optionFunc(func(o *options) { o.query["book_id"] = bookID })
}

// WithCreatedOn created_on获取
func (obj *_BookTitleMgr) WithCreatedOn(createdOn time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_on"] = createdOn })
}

// WithUpdatedOn updated_on获取
func (obj *_BookTitleMgr) WithUpdatedOn(updatedOn time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_on"] = updatedOn })
}

// GetByOption 功能选项模式获取
func (obj *_BookTitleMgr) GetByOption(opts ...Option) (result BookTitle, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.Table(obj.GetTableName()).Where(options.query).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("book_book").Where("id = ?", result.BookID).Find(&result.BookBook).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_BookTitleMgr) GetByOptions(opts ...Option) (results []*BookTitle, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.Table(obj.GetTableName()).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("book_book").Where("id = ?", results[i].BookID).Find(&results[i].BookBook).Error; err != nil { //
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
func (obj *_BookTitleMgr) GetFromID(id int64) (result BookTitle, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("book_book").Where("id = ?", result.BookID).Find(&result.BookBook).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromID 批量查找
func (obj *_BookTitleMgr) GetBatchFromID(ids []int64) (results []*BookTitle, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("book_book").Where("id = ?", results[i].BookID).Find(&results[i].BookBook).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromTitleName 通过title_name获取内容
func (obj *_BookTitleMgr) GetFromTitleName(titleName string) (results []*BookTitle, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("title_name = ?", titleName).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("book_book").Where("id = ?", results[i].BookID).Find(&results[i].BookBook).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromTitleName 批量查找
func (obj *_BookTitleMgr) GetBatchFromTitleName(titleNames []string) (results []*BookTitle, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("title_name IN (?)", titleNames).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("book_book").Where("id = ?", results[i].BookID).Find(&results[i].BookBook).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromSort 通过sort获取内容
func (obj *_BookTitleMgr) GetFromSort(sort int) (results []*BookTitle, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sort = ?", sort).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("book_book").Where("id = ?", results[i].BookID).Find(&results[i].BookBook).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromSort 批量查找
func (obj *_BookTitleMgr) GetBatchFromSort(sorts []int) (results []*BookTitle, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sort IN (?)", sorts).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("book_book").Where("id = ?", results[i].BookID).Find(&results[i].BookBook).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromBookID 通过book_id获取内容
func (obj *_BookTitleMgr) GetFromBookIDAndTitleName(bookID int64, titleName string) (result BookTitle, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("book_id = ? and title_name = ?", bookID, titleName).Take(&result).Error
	return
}

// GetFromBookID 通过book_id获取内容
func (obj *_BookTitleMgr) GetMaxSortFromBookID(bookID int64) (result BookTitle, err error) {
	err = obj.DB.Table(obj.GetTableName()).Select("max(sort)").Where("book_id = ?", bookID).Take(result).Error
	if err == nil && obj.isRelated {
		panic(err)
	}
	return
}

// GetBatchFromBookID 批量查找
func (obj *_BookTitleMgr) GetBatchFromBookID(bookIDs []int64) (results []*BookTitle, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("book_id IN (?)", bookIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("book_book").Where("id = ?", results[i].BookID).Find(&results[i].BookBook).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreatedOn 通过created_on获取内容
func (obj *_BookTitleMgr) GetFromCreatedOn(createdOn time.Time) (results []*BookTitle, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("created_on = ?", createdOn).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("book_book").Where("id = ?", results[i].BookID).Find(&results[i].BookBook).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreatedOn 批量查找
func (obj *_BookTitleMgr) GetBatchFromCreatedOn(createdOns []time.Time) (results []*BookTitle, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("created_on IN (?)", createdOns).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("book_book").Where("id = ?", results[i].BookID).Find(&results[i].BookBook).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdatedOn 通过updated_on获取内容
func (obj *_BookTitleMgr) GetFromUpdatedOn(updatedOn time.Time) (results []*BookTitle, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("updated_on = ?", updatedOn).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("book_book").Where("id = ?", results[i].BookID).Find(&results[i].BookBook).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdatedOn 批量查找
func (obj *_BookTitleMgr) GetBatchFromUpdatedOn(updatedOns []time.Time) (results []*BookTitle, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("updated_on IN (?)", updatedOns).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("book_book").Where("id = ?", results[i].BookID).Find(&results[i].BookBook).Error; err != nil { //
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
func (obj *_BookTitleMgr) FetchByPrimaryKey(id int64) (result BookTitle, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.New().Table("book_book").Where("id = ?", result.BookID).Find(&result.BookBook).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByBookBookTitleBookBookID  获取多个内容
func (obj *_BookTitleMgr) FetchIndexByBookBookTitleBookBookID(bookID int64) (results []*BookTitle, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("book_id = ?", bookID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.New().Table("book_book").Where("id = ?", results[i].BookID).Find(&results[i].BookBook).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
