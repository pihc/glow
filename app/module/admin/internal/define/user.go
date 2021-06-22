package define

import (
	"github.com/gogf/gf/os/gtime"
	"payget/app/model"
)

type RspUserInfo struct {
	Id           uint        `orm:"id,primary"    json:"id"`           // 主键ID
	Realname     string      `orm:"realname"      json:"realname"`     // 真实姓名
	Nickname     string      `orm:"nickname"      json:"nickname"`     // 昵称
	Gender       uint        `orm:"gender"        json:"gender"`       // 性别:1男 2女 3保密
	Avatar       string      `orm:"avatar"        json:"avatar"`       // 头像
	Mobile       string      `orm:"mobile"        json:"mobile"`       // 手机号码
	Email        string      `orm:"email"         json:"email"`        // 邮箱地址
	Birthday     *gtime.Time `orm:"birthday"      json:"birthday"`     // 出生日期
	DeptId       uint        `orm:"dept_id"       json:"deptId"`       // 部门ID
	LevelId      uint        `orm:"level_id"      json:"levelId"`      // 职级ID
	PositionId   uint        `orm:"position_id"   json:"positionId"`   // 岗位ID
	ProvinceCode string      `orm:"province_code" json:"provinceCode"` // 省份ID
	CityCode     string      `orm:"city_code"     json:"cityCode"`     // 市区ID
	DistrictCode string      `orm:"district_code" json:"districtCode"` // 区县ID
	Address      string      `orm:"address"       json:"address"`      // 详细地址
	CityName     string      `orm:"city_name"     json:"cityName"`     // 所属城市
	Username     string      `orm:"username"      json:"username"`     // 登录用户名
	Password     string      `orm:"password"      json:"password"`     // 登录密码
	Salt         string      `orm:"salt"          json:"salt"`         // 盐加密
	Intro        string      `orm:"intro"         json:"intro"`        // 个人简介
	Status       uint        `orm:"status"        json:"status"`       // 状态：1正常 2禁用
	Note         string      `orm:"note"          json:"note"`         // 备注
	Sort         uint        `orm:"sort"          json:"sort"`         // 显示顺序
	LoginNum     uint        `orm:"login_num"     json:"loginNum"`     // 登录次数
	LoginIp      string      `orm:"login_ip"      json:"loginIp"`      // 最近登录IP
	LoginTime    *gtime.Time `orm:"login_time"    json:"loginTime"`    // 最近登录时间
	CreateUser   uint        `orm:"create_user"   json:"createUser"`   // 添加人
	CreateTime   *gtime.Time `orm:"create_time"   json:"createTime"`   // 创建时间
	UpdateUser   uint        `orm:"update_user"   json:"updateUser"`   // 更新人
	UpdateTime   *gtime.Time `orm:"update_time"   json:"updateTime"`   // 更新时间
	Mark         uint        `orm:"mark"          json:"mark"`         // 有效标识(1正常 0删除)
	/**
	 * 角色列表
	 */
	Roles          []*model.SysRole `json:"roles"`
	Authorities    []*model.SysMenu `json:"authorities"`
	PermissionList []string         `json:"permissionList"`
}
