package ormhook

import (
	"github.com/sunweiwe/paddle/core/common"
	"gorm.io/gorm"
)

const (
	_createdBy = "created_by"
	_updatedBy = "updated_by"
)

func setCreatedByUpdatedByForCreateHook(db *gorm.DB) {
	currentUser, err := common.UserFromContext(db.Statement.Context)
	if err != nil {
		return
	}
	field := db.Statement.Schema.LookUpField(_createdBy)
	if field != nil {
		db.Statement.SetColumn(_createdBy, currentUser.GetID(), true)
	}

	field = db.Statement.Schema.LookUpField(_updatedBy)
	if field != nil {
		db.Statement.SetColumn(_updatedBy, currentUser.GetID(), true)
	}
}

func setUpdatedByForUpdateDeleteHook(db *gorm.DB) {
	currentUser, err := common.UserFromContext(db.Statement.Context)
	if err != nil {
		return
	}
	field := db.Statement.Schema.LookUpField(_updatedBy)
	if field != nil {
		db.Statement.SetColumn(_updatedBy, currentUser.GetID())
	}
}

func RegisterCustomHooks(db *gorm.DB) {
	_ = db.Callback().Create().After("gorm:before_create").Before("gorm:create").
		Register("add_create_by", setCreatedByUpdatedByForCreateHook)

	_ = db.Callback().Update().After("gorm:before_update").Before("gorm:update").
		Register("add_updated_by", setUpdatedByForUpdateDeleteHook)

	_ = db.Callback().Delete().After("gorm:before_delete").Before("gorm:delete").
		Register("add_update_by", setUpdatedByForUpdateDeleteHook)

}
