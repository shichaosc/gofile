package model

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type _UserAutherMgr struct {
	*_BaseMgr
}

// UserAutherMgr open func
func UserAutherMgr(db *gorm.DB) *_UserAutherMgr {
	if db == nil {
		panic(fmt.Errorf("UserAutherMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UserAutherMgr{_BaseMgr: &_BaseMgr{DB: db.Table("user_auther"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UserAutherMgr) GetTableName() string {
	return "user_auther"
}

// Get 获取
func (obj *_UserAutherMgr) Get() (result UserAuther, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_UserAutherMgr) Gets() (results []*UserAuther, err error) {
	err = obj.DB.Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 id
func (obj *_UserAutherMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUsername username获取
func (obj *_UserAutherMgr) WithUsername(username string) Option {
	return optionFunc(func(o *options) { o.query["username"] = username })
}

// WithRealname realname获取
func (obj *_UserAutherMgr) WithRealname(realname string) Option {
	return optionFunc(func(o *options) { o.query["realname"] = realname })
}

// WithIntroduction introduction获取
func (obj *_UserAutherMgr) WithIntroduction(introduction string) Option {
	return optionFunc(func(o *options) { o.query["introduction"] = introduction })
}

// WithAvatar avatar获取
func (obj *_UserAutherMgr) WithAvatar(avatar string) Option {
	return optionFunc(func(o *options) { o.query["avatar"] = avatar })
}

// WithGender gender获取
func (obj *_UserAutherMgr) WithGender(gender int) Option {
	return optionFunc(func(o *options) { o.query["gender"] = gender })
}

// WithOnDelete on_delete获取
func (obj *_UserAutherMgr) WithOnDelete(onDelete int) Option {
	return optionFunc(func(o *options) { o.query["on_delete"] = onDelete })
}

// WithCreatedOn created_on获取
func (obj *_UserAutherMgr) WithCreatedOn(createdOn time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_on"] = createdOn })
}

// WithUpdatedOn updated_on获取
func (obj *_UserAutherMgr) WithUpdatedOn(updatedOn time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_on"] = updatedOn })
}

// GetByOption 功能选项模式获取
func (obj *_UserAutherMgr) GetByOption(opts ...Option) (result UserAuther, err error) {
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
func (obj *_UserAutherMgr) GetByOptions(opts ...Option) (results []*UserAuther, err error) {
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

// GetFromID 通过id获取内容 id
func (obj *_UserAutherMgr) GetFromID(id int) (result UserAuther, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找 id
func (obj *_UserAutherMgr) GetBatchFromID(ids []int) (results []*UserAuther, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromUsername 通过username获取内容
func (obj *_UserAutherMgr) GetFromUsername(username string) (result UserAuther, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("username = ?", username).First(&result).Error

	return
}

// GetBatchFromUsername 批量查找
func (obj *_UserAutherMgr) GetBatchFromUsername(usernames []string) (results []*UserAuther, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("username IN (?)", usernames).Find(&results).Error

	return
}

// GetFromRealname 通过realname获取内容
func (obj *_UserAutherMgr) GetFromRealname(realname string) (results []*UserAuther, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("realname = ?", realname).Find(&results).Error

	return
}

// GetBatchFromRealname 批量查找
func (obj *_UserAutherMgr) GetBatchFromRealname(realnames []string) (results []*UserAuther, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("realname IN (?)", realnames).Find(&results).Error

	return
}

// GetFromIntroduction 通过introduction获取内容
func (obj *_UserAutherMgr) GetFromIntroduction(introduction string) (results []*UserAuther, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("introduction = ?", introduction).Find(&results).Error

	return
}

// GetBatchFromIntroduction 批量查找
func (obj *_UserAutherMgr) GetBatchFromIntroduction(introductions []string) (results []*UserAuther, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("introduction IN (?)", introductions).Find(&results).Error

	return
}

// GetFromAvatar 通过avatar获取内容
func (obj *_UserAutherMgr) GetFromAvatar(avatar string) (results []*UserAuther, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("avatar = ?", avatar).Find(&results).Error

	return
}

// GetBatchFromAvatar 批量查找
func (obj *_UserAutherMgr) GetBatchFromAvatar(avatars []string) (results []*UserAuther, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("avatar IN (?)", avatars).Find(&results).Error

	return
}

// GetFromGender 通过gender获取内容
func (obj *_UserAutherMgr) GetFromGender(gender int) (results []*UserAuther, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("gender = ?", gender).Find(&results).Error

	return
}

// GetBatchFromGender 批量查找
func (obj *_UserAutherMgr) GetBatchFromGender(genders []int) (results []*UserAuther, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("gender IN (?)", genders).Find(&results).Error

	return
}

// GetFromOnDelete 通过on_delete获取内容
func (obj *_UserAutherMgr) GetFromOnDelete(onDelete int) (results []*UserAuther, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("on_delete = ?", onDelete).Find(&results).Error

	return
}

// GetBatchFromOnDelete 批量查找
func (obj *_UserAutherMgr) GetBatchFromOnDelete(onDeletes []int) (results []*UserAuther, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("on_delete IN (?)", onDeletes).Find(&results).Error

	return
}

// GetFromCreatedOn 通过created_on获取内容
func (obj *_UserAutherMgr) GetFromCreatedOn(createdOn time.Time) (results []*UserAuther, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("created_on = ?", createdOn).Find(&results).Error

	return
}

// GetBatchFromCreatedOn 批量查找
func (obj *_UserAutherMgr) GetBatchFromCreatedOn(createdOns []time.Time) (results []*UserAuther, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("created_on IN (?)", createdOns).Find(&results).Error

	return
}

// GetFromUpdatedOn 通过updated_on获取内容
func (obj *_UserAutherMgr) GetFromUpdatedOn(updatedOn time.Time) (results []*UserAuther, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("updated_on = ?", updatedOn).Find(&results).Error

	return
}

// GetBatchFromUpdatedOn 批量查找
func (obj *_UserAutherMgr) GetBatchFromUpdatedOn(updatedOns []time.Time) (results []*UserAuther, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("updated_on IN (?)", updatedOns).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_UserAutherMgr) FetchByPrimaryKey(id int) (result UserAuther, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByGender  获取多个内容
func (obj *_UserAutherMgr) FetchIndexByGender(gender int) (results []*UserAuther, err error) {
	err = obj.DB.Table(obj.GetTableName()).Where("gender = ?", gender).Find(&results).Error

	return
}
