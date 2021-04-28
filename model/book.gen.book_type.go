package model

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _BookTypeMgr struct {
	*_BaseMgr
}

// BookTypeMgr open func
func BookTypeMgr(db *gorm.DB) *_BookTypeMgr {
	if db == nil {
		panic(fmt.Errorf("BookTypeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_BookTypeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("book_type"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_BookTypeMgr) GetTableName() string {
	return "book_type"
}

// Get 获取
func (obj *_BookTypeMgr) Get() (result BookType, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_BookTypeMgr) Gets() (results []*BookType, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_BookTypeMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithTypeName type_name获取
func (obj *_BookTypeMgr) WithTypeName(typeName string) Option {
	return optionFunc(func(o *options) { o.query["type_name"] = typeName })
}

// WithSort sort获取
func (obj *_BookTypeMgr) WithSort(sort int) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// WithCreatedOn created_on获取
func (obj *_BookTypeMgr) WithCreatedOn(createdOn time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_on"] = createdOn })
}

// WithUpdatedOn updated_on获取
func (obj *_BookTypeMgr) WithUpdatedOn(updatedOn time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_on"] = updatedOn })
}

// GetByOption 功能选项模式获取
func (obj *_BookTypeMgr) GetByOption(opts ...Option) (result BookType, err error) {
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
func (obj *_BookTypeMgr) GetByOptions(opts ...Option) (results []*BookType, err error) {
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
func (obj *_BookTypeMgr) GetFromID(id int) (result BookType, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_BookTypeMgr) GetBatchFromID(ids []int) (results []*BookType, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromTypeName 通过type_name获取内容
func (obj *_BookTypeMgr) GetFromTypeName(typeName string) (result BookType, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("type_name = ?", typeName).Take(&result).Error

	return
}

// GetBatchFromTypeName 批量查找
func (obj *_BookTypeMgr) GetBatchFromTypeName(typeNames []string) (results []*BookType, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("type_name IN (?)", typeNames).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容
func (obj *_BookTypeMgr) GetFromSort(sort int) (results []*BookType, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sort = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量查找
func (obj *_BookTypeMgr) GetBatchFromSort(sorts []int) (results []*BookType, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("sort IN (?)", sorts).Find(&results).Error

	return
}

// GetFromCreatedOn 通过created_on获取内容
func (obj *_BookTypeMgr) GetFromCreatedOn(createdOn time.Time) (results []*BookType, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("created_on = ?", createdOn).Find(&results).Error

	return
}

// GetBatchFromCreatedOn 批量查找
func (obj *_BookTypeMgr) GetBatchFromCreatedOn(createdOns []time.Time) (results []*BookType, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("created_on IN (?)", createdOns).Find(&results).Error

	return
}

// GetFromUpdatedOn 通过updated_on获取内容
func (obj *_BookTypeMgr) GetFromUpdatedOn(updatedOn time.Time) (results []*BookType, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("updated_on = ?", updatedOn).Find(&results).Error

	return
}

// GetBatchFromUpdatedOn 批量查找
func (obj *_BookTypeMgr) GetBatchFromUpdatedOn(updatedOns []time.Time) (results []*BookType, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("updated_on IN (?)", updatedOns).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_BookTypeMgr) FetchByPrimaryKey(id int) (result BookType, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
