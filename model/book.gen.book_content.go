package model

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _BookContentMgr struct {
	*_BaseMgr
}

// BookContentMgr open func
func BookContentMgr(db *gorm.DB) *_BookContentMgr {
	if db == nil {
		panic(fmt.Errorf("BookContentMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_BookContentMgr{_BaseMgr: &_BaseMgr{DB: db.Table("book_content"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_BookContentMgr) GetTableName() string {
	return "book_content"
}

// Get 获取
func (obj *_BookContentMgr) Get() (result BookContent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_BookContentMgr) Gets() (results []*BookContent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_BookContentMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithContent content获取
func (obj *_BookContentMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// WithTitleID title_id获取
func (obj *_BookContentMgr) WithTitleID(titleID int64) Option {
	return optionFunc(func(o *options) { o.query["title_id"] = titleID })
}

// WithCreatedOn created_on获取
func (obj *_BookContentMgr) WithCreatedOn(createdOn time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_on"] = createdOn })
}

// WithUpdatedOn updated_on获取
func (obj *_BookContentMgr) WithUpdatedOn(updatedOn time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_on"] = updatedOn })
}

// GetByOption 功能选项模式获取
func (obj *_BookContentMgr) GetByOption(opts ...Option) (result BookContent, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_BookContentMgr) GetByOptions(opts ...Option) (results []*BookContent, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_BookContentMgr) GetFromID(id int64) (result BookContent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_BookContentMgr) GetBatchFromID(ids []int64) (results []*BookContent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromContent 通过content获取内容
func (obj *_BookContentMgr) GetFromContent(content string) (results []*BookContent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("content = ?", content).Find(&results).Error

	return
}

// GetBatchFromContent 批量查找
func (obj *_BookContentMgr) GetBatchFromContent(contents []string) (results []*BookContent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("content IN (?)", contents).Find(&results).Error

	return
}

// GetFromTitleID 通过title_id获取内容
func (obj *_BookContentMgr) GetFromTitleID(titleID int64) (results []*BookContent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("title_id = ?", titleID).Find(&results).Error

	return
}

// GetBatchFromTitleID 批量查找
func (obj *_BookContentMgr) GetBatchFromTitleID(titleIDs []int64) (results []*BookContent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("title_id IN (?)", titleIDs).Find(&results).Error

	return
}

// GetFromCreatedOn 通过created_on获取内容
func (obj *_BookContentMgr) GetFromCreatedOn(createdOn time.Time) (results []*BookContent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("created_on = ?", createdOn).Find(&results).Error

	return
}

// GetBatchFromCreatedOn 批量查找
func (obj *_BookContentMgr) GetBatchFromCreatedOn(createdOns []time.Time) (results []*BookContent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("created_on IN (?)", createdOns).Find(&results).Error

	return
}

// GetFromUpdatedOn 通过updated_on获取内容
func (obj *_BookContentMgr) GetFromUpdatedOn(updatedOn time.Time) (results []*BookContent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("updated_on = ?", updatedOn).Find(&results).Error

	return
}

// GetBatchFromUpdatedOn 批量查找
func (obj *_BookContentMgr) GetBatchFromUpdatedOn(updatedOns []time.Time) (results []*BookContent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("updated_on IN (?)", updatedOns).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_BookContentMgr) FetchByPrimaryKey(id int64) (result BookContent, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
