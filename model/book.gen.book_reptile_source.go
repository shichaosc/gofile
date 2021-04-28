package model

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _BookReptileSourceMgr struct {
	*_BaseMgr
}

// BookReptileSourceMgr open func
func BookReptileSourceMgr(db *gorm.DB) *_BookReptileSourceMgr {
	if db == nil {
		panic(fmt.Errorf("BookReptileSourceMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_BookReptileSourceMgr{_BaseMgr: &_BaseMgr{DB: db.Table("book_reptile_source"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_BookReptileSourceMgr) GetTableName() string {
	return "book_reptile_source"
}

// Get 获取
func (obj *_BookReptileSourceMgr) Get() (result BookReptileSource, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_BookReptileSourceMgr) Gets() (results []*BookReptileSource, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_BookReptileSourceMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithContentID content_id获取
func (obj *_BookReptileSourceMgr) WithContentID(contentID int64) Option {
	return optionFunc(func(o *options) { o.query["content_id"] = contentID })
}

// WithSource source获取
func (obj *_BookReptileSourceMgr) WithSource(source string) Option {
	return optionFunc(func(o *options) { o.query["source"] = source })
}

// WithSourceURL source_url获取 url
func (obj *_BookReptileSourceMgr) WithSourceURL(sourceURL string) Option {
	return optionFunc(func(o *options) { o.query["source_url"] = sourceURL })
}

// WithCreatedOn created_on获取
func (obj *_BookReptileSourceMgr) WithCreatedOn(createdOn time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_on"] = createdOn })
}

// WithUpdatedOn updated_on获取
func (obj *_BookReptileSourceMgr) WithUpdatedOn(updatedOn time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_on"] = updatedOn })
}

// GetByOption 功能选项模式获取
func (obj *_BookReptileSourceMgr) GetByOption(opts ...Option) (result BookReptileSource, err error) {
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
func (obj *_BookReptileSourceMgr) GetByOptions(opts ...Option) (results []*BookReptileSource, err error) {
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
func (obj *_BookReptileSourceMgr) GetFromID(id int64) (result BookReptileSource, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_BookReptileSourceMgr) GetBatchFromID(ids []int64) (results []*BookReptileSource, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromContentID 通过content_id获取内容
func (obj *_BookReptileSourceMgr) GetFromContentID(contentID int64) (results []*BookReptileSource, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("content_id = ?", contentID).Find(&results).Error

	return
}

// GetBatchFromContentID 批量查找
func (obj *_BookReptileSourceMgr) GetBatchFromContentID(contentIDs []int64) (results []*BookReptileSource, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("content_id IN (?)", contentIDs).Find(&results).Error

	return
}

// GetFromSource 通过source获取内容
func (obj *_BookReptileSourceMgr) GetFromSource(source string) (results []*BookReptileSource, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("source = ?", source).Find(&results).Error

	return
}

// GetBatchFromSource 批量查找
func (obj *_BookReptileSourceMgr) GetBatchFromSource(sources []string) (results []*BookReptileSource, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("source IN (?)", sources).Find(&results).Error

	return
}

// GetFromSourceURL 通过source_url获取内容 url
func (obj *_BookReptileSourceMgr) GetFromSourceURL(sourceURL string) (results []*BookReptileSource, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("source_url = ?", sourceURL).Find(&results).Error

	return
}

// GetBatchFromSourceURL 批量查找 url
func (obj *_BookReptileSourceMgr) GetBatchFromSourceURL(sourceURLs []string) (results []*BookReptileSource, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("source_url IN (?)", sourceURLs).Find(&results).Error

	return
}

// GetFromCreatedOn 通过created_on获取内容
func (obj *_BookReptileSourceMgr) GetFromCreatedOn(createdOn time.Time) (results []*BookReptileSource, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("created_on = ?", createdOn).Find(&results).Error

	return
}

// GetBatchFromCreatedOn 批量查找
func (obj *_BookReptileSourceMgr) GetBatchFromCreatedOn(createdOns []time.Time) (results []*BookReptileSource, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("created_on IN (?)", createdOns).Find(&results).Error

	return
}

// GetFromUpdatedOn 通过updated_on获取内容
func (obj *_BookReptileSourceMgr) GetFromUpdatedOn(updatedOn time.Time) (results []*BookReptileSource, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("updated_on = ?", updatedOn).Find(&results).Error

	return
}

// GetBatchFromUpdatedOn 批量查找
func (obj *_BookReptileSourceMgr) GetBatchFromUpdatedOn(updatedOns []time.Time) (results []*BookReptileSource, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("updated_on IN (?)", updatedOns).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_BookReptileSourceMgr) FetchByPrimaryKey(id int64) (result BookReptileSource, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
