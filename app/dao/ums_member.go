// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"payget/app/dao/internal"
)

// umsMemberDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type umsMemberDao struct {
	*internal.UmsMemberDao
}

var (
	// UmsMember is globally public accessible object for table ums_member operations.
	UmsMember umsMemberDao
)

func init() {
	UmsMember = umsMemberDao{
		internal.NewUmsMemberDao(),
	}
}

// Fill with you ideas below.
