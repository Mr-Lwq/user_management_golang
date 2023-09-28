package dao

import (
	"fmt"
	"testing"
	"time"
	"user_management_golang/src"
)

// 创建创建表和默认用户创建
func TestNewMysql(t *testing.T) {
	cfg := MysqlCfg{
		"root",
		"Zkyy2021",
		"10.1.36.245",
		"3307",
		"user_management"}
	mysql, err := NewMysql(cfg)
	if err != nil {
		fmt.Println(err)
	}
	mysql.init()
}

// 测试创建表功能
func TestMysql_CreateTable(t *testing.T) {
	cfg := MysqlCfg{
		"root",
		"Zkyy2021",
		"10.1.36.245",
		"3307",
		"user_management"}
	mysql, _ := NewMysql(cfg)
	tables := []string{
		mysql.UserGroupTable,
		mysql.RoleTable,
		mysql.AccountTable,
	}
	for _, table := range tables {
		mysql.CreateTable(table)
	}
}

// 测试Insert 的功能
func TestMysql_Insert(t *testing.T) {
	cfg := MysqlCfg{
		"root",
		"Zkyy2021",
		"10.1.36.245",
		"3307",
		"user_management"}
	mysql, _ := NewMysql(cfg)
	admin := src.Account{
		UserId:         "admin1",
		Username:       "admin1",
		Password:       "88888888",
		Email:          "",
		Phone:          "",
		FullName:       "超级管理员",
		Roles:          []string{"administrators"},
		Status:         "activate",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		LastLoginAt:    time.Now(),
		SessionToken:   "",
		ProfilePicture: "",
		UserGroups:     []string{"administrators"},
	}
	err := mysql.Insert(admin)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("插入数据成功")
	}
}

// 测试更新数据
func TestMysql_Update(t *testing.T) {
	cfg := MysqlCfg{
		"root",
		"Zkyy2021",
		"10.1.36.245",
		"3307",
		"user_management"}
	mysql, _ := NewMysql(cfg)
	admin := src.Account{
		UserId:         "admin1",
		Username:       "admin1",
		Password:       "88888",
		Email:          "",
		Phone:          "",
		FullName:       "超级管理员",
		Roles:          []string{"administrators"},
		Status:         "",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		LastLoginAt:    time.Now(),
		SessionToken:   "",
		ProfilePicture: "",
		UserGroups:     []string{"administrators"},
	}
	err := mysql.Update(admin)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("数据更新成功")
	}
}

// 测试删除数据
func TestMysql_Del(t *testing.T) {
	cfg := MysqlCfg{
		"root",
		"Zkyy2021",
		"10.1.36.245",
		"3307",
		"user_management"}
	mysql, _ := NewMysql(cfg)
	admin := src.Account{
		UserId:         "admin1",
		Username:       "admin1",
		Password:       "88888",
		Email:          "",
		Phone:          "",
		FullName:       "超级管理员",
		Roles:          []string{"administrators"},
		Status:         "",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		LastLoginAt:    time.Now(),
		SessionToken:   "",
		ProfilePicture: "",
		UserGroups:     []string{"administrators"},
	}
	err := mysql.Del(admin)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("数据删除成功")
	}
}

// 测试查询功能
func TestMysql_Search(t *testing.T) {
	cfg := MysqlCfg{
		"root",
		"Zkyy2021",
		"10.1.36.245",
		"3307",
		"user_management"}
	mysql, _ := NewMysql(cfg)
	//admin := src.Account{
	//	UserId:         "admin",
	//	Username:       "admin",
	//	Password:       "88888",
	//	Email:          "",
	//	Phone:          "",
	//	FullName:       "超级管理员",
	//	Roles:          []string{"administrators"},
	//	Status:         "",
	//	CreatedAt:      time.Now(),
	//	UpdatedAt:      time.Now(),
	//	LastLoginAt:    time.Now(),
	//	SessionToken:   "",
	//	ProfilePicture: "",
	//	UserGroups:     []string{"administrators"},
	//}
	//role := src.Role{
	//	RoleId: "admin",
	//}
	group := src.UserGroup{
		GroupId: "administrators",
	}
	result, err := mysql.Search(group)
	if err != nil {
		fmt.Println(err)
	} else {
		// 检查查询结果是否与预期相符
		retrievedUser, ok := result.(src.UserGroup)
		if !ok {
			t.Fatalf("Expected result type: *src.Account, got: %T", result)
		}

		// 检查其他字段...
		fmt.Printf("Retrieved User: %+v\n", retrievedUser)
	}
}
