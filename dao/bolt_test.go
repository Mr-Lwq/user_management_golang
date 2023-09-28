package dao

import (
	"fmt"
	"testing"
	"time"
	"user_management_golang/src"
)

// Test1 【嵌入式键值存储数据库】能否正常创建
func TestNewMyBolt(t *testing.T) {
	mybolt, err := NewMyBolt()
	if err != nil {
		fmt.Println(err, "连接失败！")
		return
	}
	//err = db.CreateTable("user_table")
	//err = db.InsertData("user_table", []byte("key1"), []byte("value1"))
	mybolt.Close()
}

// Test2 【BoltDB】 插入数据测试
func TestMyBolt_Insert(t *testing.T) {
	mybolt, err := NewMyBolt()
	if err != nil {
		fmt.Println(err, "连接失败！")
		return
	}
	user := src.Account{
		UserId:         "user",
		Username:       "user",
		Password:       "88888888",
		Email:          "",
		Phone:          "",
		FullName:       "普通用户",
		Roles:          []string{"users"},
		Status:         "activate",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		LastLoginAt:    time.Now(),
		SessionToken:   "",
		ProfilePicture: "",
		UserGroups:     []string{""},
	}
	err = mybolt.Insert(user)
	if err != nil {
		fmt.Println(err)
	}
}

// Test3 打印指定表格里的所有记录
func TestMyBolt_PrintAll(t *testing.T) {
	mybolt, err := NewMyBolt()
	if err != nil {
		fmt.Println(err, "连接失败！")
		return
	}
	mybolt.PrintAll(mybolt.AccountTable)
}

// Test4 更新数据功能
func TestMyBolt_Update(t *testing.T) {
	mybolt, err := NewMyBolt()
	if err != nil {
		fmt.Println(err, "连接失败！")
		return
	}
	user := src.Account{
		UserId:         "user",
		Username:       "user",
		Password:       "88888",
		Email:          "",
		Phone:          "",
		FullName:       "普通用户",
		Roles:          []string{"users"},
		Status:         "activate",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		LastLoginAt:    time.Now(),
		SessionToken:   "",
		ProfilePicture: "",
		UserGroups:     []string{""},
	}
	err = mybolt.Update(user)
	if err != nil {
		fmt.Println(err)
	}
}

// Test5 删除数据
func TestMyBolt_Del(t *testing.T) {
	mybolt, err := NewMyBolt()
	if err != nil {
		fmt.Println(err, "连接失败！")
		return
	}
	user := src.Account{
		UserId:         "user",
		Username:       "user",
		Password:       "",
		Email:          "",
		Phone:          "",
		FullName:       "",
		Roles:          []string{"users"},
		Status:         "",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		LastLoginAt:    time.Now(),
		SessionToken:   "",
		ProfilePicture: "",
		UserGroups:     []string{""},
	}
	err = mybolt.Del(user)
	if err != nil {
		fmt.Println(err)
	}
}

// Test6 查询数据
func TestMyBolt_Search(t *testing.T) {
	mybolt, err := NewMyBolt()
	if err != nil {
		fmt.Println(err, "连接失败！")
		return
	}
	user := src.Account{
		UserId:         "admin",
		Username:       "admin",
		Password:       "",
		Email:          "",
		Phone:          "",
		FullName:       "",
		Roles:          []string{"users"},
		Status:         "",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		LastLoginAt:    time.Now(),
		SessionToken:   "",
		ProfilePicture: "",
		UserGroups:     []string{""},
	}
	// 查询数据
	result, err := mybolt.Search(user)
	if err != nil {
		t.Fatalf("Error searching data: %v", err)
	}

	// 检查查询结果是否与预期相符
	retrievedUser, ok := result.(*src.Account)
	if !ok {
		t.Fatalf("Expected result type: *src.Account, got: %T", result)
	}

	// 这里可以检查结构体字段的值是否符合预期
	if retrievedUser.UserId != user.UserId {
		t.Errorf("Expected UserId: %s, got: %s", user.UserId, retrievedUser.UserId)
	}
	// 检查其他字段...

	fmt.Printf("Retrieved User: %+v\n", retrievedUser)
}
