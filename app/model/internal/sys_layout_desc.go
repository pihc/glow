// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// SysLayoutDesc is the golang structure for table sys_layout_desc.
type SysLayoutDesc struct {
	Id         uint        `orm:"id,primary"  json:"id"`         // 主键ID
	LocDesc    string      `orm:"loc_desc"    json:"locDesc"`    // 页面位置描述
	LocId      uint        `orm:"loc_id"      json:"locId"`      // 位置编号
	ItemId     uint        `orm:"item_id"     json:"itemId"`     // 站点ID
	Sort       uint        `orm:"sort"        json:"sort"`       // 排序
	CreateUser uint        `orm:"create_user" json:"createUser"` // 添加人
	CreateTime *gtime.Time `orm:"create_time" json:"createTime"` // 添加时间
	UpdateUser uint        `orm:"update_user" json:"updateUser"` // 更新人
	UpdateTime *gtime.Time `orm:"update_time" json:"updateTime"` // 更新时间
	Mark       uint        `orm:"mark"        json:"mark"`       // 有效标识(1正常 0删除)
}
