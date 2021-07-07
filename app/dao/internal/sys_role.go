// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// SysRoleDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type SysRoleDao struct {
	gmvc.M                 // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	DB      gdb.DB         // DB is the raw underlying database management object.
	Table   string         // Table is the table name of the DAO.
	Columns sysRoleColumns // Columns contains all the columns of Table that for convenient usage.
}

// SysRoleColumns defines and stores column names for table sys_role.
type sysRoleColumns struct {
	Id        string // 主键ID
	Name      string // 角色名称
	Code      string // 角色标签
	Note      string // 备注
	Sort      string // 排序
	CreatedBy string // 添加人
	CreatedAt string // 添加时间
	UpdatedBy string // 更新人
	UpdatedAt string // 更新时间
}

func NewSysRoleDao() *SysRoleDao {
	return &SysRoleDao{
		M:     g.DB("default").Model("sys_role").Safe(),
		DB:    g.DB("default"),
		Table: "sys_role",
		Columns: sysRoleColumns{
			Id:        "id",
			Name:      "name",
			Code:      "code",
			Note:      "note",
			Sort:      "sort",
			CreatedBy: "created_by",
			CreatedAt: "created_at",
			UpdatedBy: "updated_by",
			UpdatedAt: "updated_at",
		},
	}
}
