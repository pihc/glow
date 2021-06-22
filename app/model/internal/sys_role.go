// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// SysRole is the golang structure for table sys_role.
type SysRole struct {
	Id         uint        `orm:"id,primary"  json:"id"`         // 主键ID
	Name       string      `orm:"name"        json:"name"`       // 角色名称
	Code       string      `orm:"code"        json:"code"`       // 角色标签
	Status     uint        `orm:"status"      json:"status"`     // 状态：1正常 2禁用
	Note       string      `orm:"note"        json:"note"`       // 备注
	Sort       uint        `orm:"sort"        json:"sort"`       // 排序
	CreateUser uint        `orm:"create_user" json:"createUser"` // 添加人
	CreateTime *gtime.Time `orm:"create_time" json:"createTime"` // 添加时间
	UpdateUser uint        `orm:"update_user" json:"updateUser"` // 更新人
	UpdateTime *gtime.Time `orm:"update_time" json:"updateTime"` // 更新时间
	Mark       uint        `orm:"mark"        json:"mark"`       // 有效标识(1正常 0删除)
}
