// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// SysOperLogDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type SysOperLogDao struct {
	gmvc.M                    // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	DB      gdb.DB            // DB is the raw underlying database management object.
	Table   string            // Table is the table name of the DAO.
	Columns sysOperLogColumns // Columns contains all the columns of Table that for convenient usage.
}

// SysOperLogColumns defines and stores column names for table sys_oper_log.
type sysOperLogColumns struct {
	Id            string // 主键ID
	Title         string // 日志标题
	LogType       string // 操作类型（0其它 1新增 2修改 3删除）
	OperMethod    string // 操作方法
	RequestMethod string // 请求方式
	OperType      string // 操作类型：0其他 1后台用户 2WAP用户
	OperName      string // 操作人员
	OperUrl       string // 请求URL
	OperIp        string // 主机地址
	OperLocation  string // 操作地点
	OperParam     string // 请求参数
	JsonResult    string // 返回参数
	Status        string // 操作状态（0正常 1异常）
	Note          string // 备注
	CreateUser    string // 添加人
	CreateTime    string // 操作时间
	UpdateUser    string // 更新人
	UpdateTime    string // 更新时间
	Mark          string // 有效标识
}

func NewSysOperLogDao() *SysOperLogDao {
	return &SysOperLogDao{
		M:     g.DB("default").Model("sys_oper_log").Safe(),
		DB:    g.DB("default"),
		Table: "sys_oper_log",
		Columns: sysOperLogColumns{
			Id:            "id",
			Title:         "title",
			LogType:       "log_type",
			OperMethod:    "oper_method",
			RequestMethod: "request_method",
			OperType:      "oper_type",
			OperName:      "oper_name",
			OperUrl:       "oper_url",
			OperIp:        "oper_ip",
			OperLocation:  "oper_location",
			OperParam:     "oper_param",
			JsonResult:    "json_result",
			Status:        "status",
			Note:          "note",
			CreateUser:    "create_user",
			CreateTime:    "create_time",
			UpdateUser:    "update_user",
			UpdateTime:    "update_time",
			Mark:          "mark",
		},
	}
}
