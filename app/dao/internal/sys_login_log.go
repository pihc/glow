// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// SysLoginLogDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type SysLoginLogDao struct {
	gmvc.M                     // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	DB      gdb.DB             // DB is the raw underlying database management object.
	Table   string             // Table is the table name of the DAO.
	Columns sysLoginLogColumns // Columns contains all the columns of Table that for convenient usage.
}

// SysLoginLogColumns defines and stores column names for table sys_login_log.
type sysLoginLogColumns struct {
	Id            string // 唯一性标识
	Title         string // 日志标题
	Username      string // 登录账号
	LoginIp       string // 登录IP地址
	LoginLocation string // 登录地区
	Browser       string // 浏览器类型
	Os            string // 操作系统
	Status        string // 登录状态：0成功 1失败
	Type          string // 操作类型：1登录系统 2注销系统 3刷新Token
	Note          string // 登录备注
	CreateUser    string // 添加人
	CreateTime    string // 登录时间
	UpdateUser    string // 更新人
	UpdateTime    string // 更新时间
	Mark          string // 有效标记
}

func NewSysLoginLogDao() *SysLoginLogDao {
	return &SysLoginLogDao{
		M:     g.DB("default").Model("sys_login_log").Safe(),
		DB:    g.DB("default"),
		Table: "sys_login_log",
		Columns: sysLoginLogColumns{
			Id:            "id",
			Title:         "title",
			Username:      "username",
			LoginIp:       "login_ip",
			LoginLocation: "login_location",
			Browser:       "browser",
			Os:            "os",
			Status:        "status",
			Type:          "type",
			Note:          "note",
			CreateUser:    "create_user",
			CreateTime:    "create_time",
			UpdateUser:    "update_user",
			UpdateTime:    "update_time",
			Mark:          "mark",
		},
	}
}
