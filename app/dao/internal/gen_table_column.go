// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// GenTableColumnDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type GenTableColumnDao struct {
	gmvc.M                        // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	DB      gdb.DB                // DB is the raw underlying database management object.
	Table   string                // Table is the table name of the DAO.
	Columns genTableColumnColumns // Columns contains all the columns of Table that for convenient usage.
}

// GenTableColumnColumns defines and stores column names for table gen_table_column.
type genTableColumnColumns struct {
	Id            string // 主键ID
	TableId       string // 归属表编号
	ColumnName    string // 列名称
	ColumnComment string // 列描述
	ColumnType    string // 列类型
	JavaType      string // JAVA类型
	JavaField     string // JAVA字段名
	IsPk          string // 是否主键（1是）
	IsIncrement   string // 是否自增（1是）
	IsRequired    string // 是否必填（1是）
	IsInsert      string // 是否为插入字段（1是）
	IsEdit        string // 是否编辑字段（1是）
	IsList        string // 是否列表字段（1是）
	IsQuery       string // 是否查询字段（1是）
	QueryType     string // 查询方式（等于、不等于、大于、小于、范围）
	HtmlType      string // 显示类型（文本框、文本域、下拉框、复选框、单选框、日期控件）
	DictType      string // 字典类型
	Sort          string // 排序
	CreateUser    string // 添加人
	CreateTime    string // 创建时间
	UpdateUser    string // 更新人
	UpdateTime    string // 更新时间
	Mark          string // 有效标识
}

func NewGenTableColumnDao() *GenTableColumnDao {
	return &GenTableColumnDao{
		M:     g.DB("default").Model("gen_table_column").Safe(),
		DB:    g.DB("default"),
		Table: "gen_table_column",
		Columns: genTableColumnColumns{
			Id:            "id",
			TableId:       "table_id",
			ColumnName:    "column_name",
			ColumnComment: "column_comment",
			ColumnType:    "column_type",
			JavaType:      "java_type",
			JavaField:     "java_field",
			IsPk:          "is_pk",
			IsIncrement:   "is_increment",
			IsRequired:    "is_required",
			IsInsert:      "is_insert",
			IsEdit:        "is_edit",
			IsList:        "is_list",
			IsQuery:       "is_query",
			QueryType:     "query_type",
			HtmlType:      "html_type",
			DictType:      "dict_type",
			Sort:          "sort",
			CreateUser:    "create_user",
			CreateTime:    "create_time",
			UpdateUser:    "update_user",
			UpdateTime:    "update_time",
			Mark:          "mark",
		},
	}
}
