// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// SysDictDataDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type SysDictDataDao struct {
	gmvc.M                     // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	DB      gdb.DB             // DB is the raw underlying database management object.
	Table   string             // Table is the table name of the DAO.
	Columns sysDictDataColumns // Columns contains all the columns of Table that for convenient usage.
}

// SysDictDataColumns defines and stores column names for table sys_dict_data.
type sysDictDataColumns struct {
	Id         string // 主键ID
	Name       string // 字典项名称
	Code       string // 字典项值
	DictId     string // 字典类型ID
	Status     string // 状态：1在用 2停用
	Note       string // 备注
	Sort       string // 显示顺序
	CreateUser string // 添加人
	CreateTime string // 添加时间
	UpdateUser string // 更新人
	UpdateTime string // 更新时间
	Mark       string // 有效标记
}

func NewSysDictDataDao() *SysDictDataDao {
	return &SysDictDataDao{
		M:     g.DB("default").Model("sys_dict_data").Safe(),
		DB:    g.DB("default"),
		Table: "sys_dict_data",
		Columns: sysDictDataColumns{
			Id:         "id",
			Name:       "name",
			Code:       "code",
			DictId:     "dict_id",
			Status:     "status",
			Note:       "note",
			Sort:       "sort",
			CreateUser: "create_user",
			CreateTime: "create_time",
			UpdateUser: "update_user",
			UpdateTime: "update_time",
			Mark:       "mark",
		},
	}
}
