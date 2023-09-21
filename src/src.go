package src

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

type Authenticator interface {
	Login(username, password string) (bool, error) // 登录方法
	Logout() error                                 // 注销方法
}

type BoltTable interface {
	GetId() string                             // 获取不同表的相应ID字段
	GetValue(fieldName string) ([]byte, error) // 获取不同字段的值的byte
	ToBytes() ([]byte, error)                  // 将这个结构体转换成bytes
}

type Account struct {
	UserId         string    // 用户id
	Username       string    // 用户名
	Password       string    // 密码
	Email          string    // 电子邮件地址
	Phone          string    // 电话
	FullName       string    // 全名
	Role           []string  // 角色列表
	Status         string    // 账户状态 (激活、锁定、禁用等)
	CreatedAt      time.Time // 账户创建时间
	UpdatedAt      time.Time // 账户更新时间
	LastLoginAt    time.Time // 最后登录时间
	SessionToken   string    // 用户会话令牌
	ProfilePicture string    // 个人资料图片的URL或文件路径
	UserGroupId    []string  // 用户组
}

func (a Account) Login(username, password string) (bool, error) {
	return false, nil
}

func (a Account) Logout() error {
	return nil
}

func (a Account) GetId() string {
	return a.UserId
}

func (a Account) GetValue(fieldName string) ([]byte, error) {
	v := reflect.ValueOf(a)
	field := v.FieldByName(fieldName)

	if !field.IsValid() {
		return nil, fmt.Errorf("Field not found: %service", fieldName)
	}
	// 将字段的值转换为 []byte
	valueBytes, ok := field.Interface().([]byte)
	if !ok {
		return nil, fmt.Errorf("Field value cannot be converted to []byte")
	}
	return valueBytes, nil
}

func (a Account) ToBytes() ([]byte, error) {
	data, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}
	return data, nil
}

type UserGroup struct {
	GroupId        string    // 用户组ID
	GroupName      string    // 用户组名称
	Description    string    // 用户组描述
	Permissions    []string  // 用户组权限（可以是权限名称的列表）
	Members        []string  // 用户组成员（可以是用户ID的列表）
	CreationTime   time.Time // 用户组创建时间
	LastUpdateTime time.Time // 用户组最后更新时间
}

func (ug UserGroup) GetId() string {
	return ug.GroupId
}

func (ug UserGroup) GetValue(fieldName string) ([]byte, error) {
	v := reflect.ValueOf(ug)
	field := v.FieldByName(fieldName)

	if !field.IsValid() {
		return nil, fmt.Errorf("Field not found: %service", fieldName)
	}
	// 将字段的值转换为 []byte
	valueBytes, ok := field.Interface().([]byte)
	if !ok {
		return nil, fmt.Errorf("Field value cannot be converted to []byte")
	}
	return valueBytes, nil
}

func (ug UserGroup) ToBytes() ([]byte, error) {
	data, err := json.Marshal(ug)
	if err != nil {
		return nil, err
	}
	return data, nil
}

type Role struct {
	RoleId       string    // 角色ID
	RoleName     string    // 角色名称
	Description  string    // 角色描述
	Permissions  []string  // 角色权限（可以是权限名称的列表）
	CreationTime time.Time // 角色创建时间
}

func (r Role) GetId() string {
	return r.RoleId
}

func (r Role) GetValue(fieldName string) ([]byte, error) {
	v := reflect.ValueOf(r)
	field := v.FieldByName(fieldName)

	if !field.IsValid() {
		return nil, fmt.Errorf("Field not found: %service", fieldName)
	}
	// 将字段的值转换为 []byte
	valueBytes, ok := field.Interface().([]byte)
	if !ok {
		return nil, fmt.Errorf("Field value cannot be converted to []byte")
	}
	return valueBytes, nil
}

func (r Role) ToBytes() ([]byte, error) {
	data, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	return data, nil
}
