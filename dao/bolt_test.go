package dao

import (
	"fmt"
	"testing"
	"user_management_golang/src"
)

// Test1 【嵌入式键值存储数据库】能否正常创建
func TestNewMyBolt(t *testing.T) {
	mybolt, err := NewMyBolt()
	if err != nil {
		fmt.Println(err, "连接失败！")
		return
	}
	//err = db.CreateBucket("user_table")
	//err = db.InsertData("user_table", []byte("key1"), []byte("value1"))
	mybolt.Close()
}

// Test2 【BoltDB】 正常的插入和读取
func TestMyBolt_InsertData(t *testing.T) {
	mybolt, err := NewMyBolt()
	if err != nil {
		fmt.Println(err, "连接失败！")
		return
	}

	user1 := src.Account{UserId: "user002"}
	err = mybolt.InsertData(user1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	data, err := mybolt.GetData(src.Account{UserId: "admin001"})
	if err != nil {
		fmt.Println("Error:", err)
		return
	} else {
		if account, ok := data.(*src.Account); ok {
			fmt.Println("account:", account)
		} else {
			fmt.Println(account)
			fmt.Println(ok)
		}
	}
}

// Test3 打印指定表格里的所有记录
func TestMyBolt_PrintAll(t *testing.T) {
	mybolt, err := NewMyBolt()
	if err != nil {
		fmt.Println(err, "连接失败！")
		return
	}
	mybolt.PrintAll(mybolt.RoleTable)
}
