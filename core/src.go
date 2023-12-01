package core

import (
	"time"
)

type Authenticator interface {
	Login(username, password string) (bool, error) // 登录方法
	Logout() error                                 // 注销方法
}

type TableData interface {
	GetTableType() string
}

type Account struct {
	UserId         string    // 用户id
	Username       string    // 用户名
	Password       string    // 密码
	Email          string    // 电子邮件地址
	Phone          string    // 电话
	FullName       string    // 全名
	Roles          []string  // 角色列表
	Status         string    // 账户状态 (激活、锁定、禁用等)
	CreatedAt      time.Time // 账户创建时间
	UpdatedAt      time.Time // 账户更新时间
	LastLoginAt    time.Time // 最后登录时间
	SessionToken   string    // 用户会话令牌
	ProfilePicture string    // 个人资料图片的URL或文件路径
	UserGroups     []string  // 用户组
}

func (a Account) GetTableType() string {
	return "Account"
}

type UserGroup struct {
	GroupId       string    // 用户组ID
	GroupLeads    string    // 组长ID
	GroupName     string    // 用户组名称
	Description   string    // 用户组描述
	Permissions   []string  // 用户组权限（可以是权限名称的列表）
	Members       []string  // 用户组成员（可以是用户ID的列表）
	CreatedAt     time.Time // 用户组创建时间
	LastUpdatedAt time.Time // 用户组最后更新时间
}

func (ug UserGroup) GetTableType() string {
	return "UserGroup"
}

type Role struct {
	RoleId      string    // 角色ID
	RoleName    string    // 角色名称
	Description string    // 角色描述
	Permissions []string  // 角色权限（可以是权限名称的列表）
	CreatedAt   time.Time // 角色创建时间
}

func (r Role) GetTableType() string {
	return "Role"
}
