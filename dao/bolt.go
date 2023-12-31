/*
Package dao is the Dao layer of the program.
This program uses two databases, BoltDB and MySQL.
If you want to keep it simple, starting Simple Mode will default to the embedded
database BoltDB.
*/

package dao

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/olekukonko/tablewriter"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"reflect"
	"time"
	"user_management_golang/core"
	"user_management_golang/utils"
)

var shortFormat = "2006-01-02 15:04:05"

// MyBolt 结构体包含 BoltDB 实例
type MyBolt struct {
	db             *bolt.DB
	UserGroupTable string
	RoleTable      string
	AccountTable   string
}

// NewMyBolt creates a new MyBolt instance for interacting with a BoltDB database.
// It opens the specified database file at the given path and returns a pointer
// to the MyBolt instance. If an error occurs during instance creation, it returns
// a non-nil error.
func NewMyBolt() (*MyBolt, error) {
	// Open the BoltDB database
	db, err := bolt.Open("BoltDB", 0600, nil)
	if err != nil {
		return nil, err
	}
	// Create and initialize the MyBolt instance
	myBolt := &MyBolt{
		db:             db,
		UserGroupTable: "user_group_table",
		RoleTable:      "role_table",
		AccountTable:   "account_table",
	}
	myBolt.init()
	return myBolt, nil
}

// init helps you initialize the BoltDB table the first time you open the program.
func (myBolt *MyBolt) init() {
	tables := [3]string{
		myBolt.AccountTable,
		myBolt.UserGroupTable,
		myBolt.RoleTable,
	}
	firstMake := false
	err := myBolt.db.View(func(tx *bolt.Tx) error {
		for _, table := range tables {
			bucket := tx.Bucket([]byte(table))
			if bucket == nil {
				firstMake = true
			}
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	// 创建表的初始化信息
	if firstMake {
		for _, table := range tables {
			err := myBolt.CreateTable(table)
			if err != nil {
				log.Printf("Failed to create bucket '%service': %v", table, err)
				return
			}
		}
		var password []byte
		fmt.Println("first setting beginning...")
		// 完成用户表的初始化信息创建
		admin := core.Account{
			UserId:   "admin",
			Username: "admin",
			Password: "88888888",
			Email:    "",
			Phone:    "",
			FullName: "超级管理员",
			Roles: []string{
				"administrators",
			},
			Status:         "offline",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
			LastLoginAt:    time.Now(),
			SessionToken:   "",
			ProfilePicture: "",
			UserGroups: []string{
				"administrators",
			},
		}
		password = []byte(admin.Password)
		hashedPassword, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
		admin.Password = string(hashedPassword)
		sessionToken, _ := utils.GenerateSessionToken(admin.UserId)
		admin.SessionToken = sessionToken
		err = myBolt.Insert(admin)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("[UserId: admin, Password: 88888888] " + "Create account table success. Add root account.")
		// 完成用户组表的初始化信息创建
		userGroup := core.UserGroup{
			GroupId:       "administrators",
			GroupLeads:    "admin",
			GroupName:     "administrators",
			Description:   "",
			Permissions:   []string{},
			Members:       []string{"admin"},
			CreatedAt:     time.Now(),
			LastUpdatedAt: time.Now(),
		}
		err = myBolt.Insert(userGroup)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("[GroupId: group, GroupName: administrators] " + "Create user_group table success. Add administrator group.")
		// 完成角色表的初始化信息创建
		role := core.Role{
			RoleId:      "admin",
			RoleName:    "admin",
			Description: "",
			Permissions: []string{},
			CreatedAt:   time.Now(),
		}
		err = myBolt.Insert(role)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("[RoleId: role, RoleName: admin] " + "Create role table success. Add admin role.")
		fmt.Println("The first setting is successful. Please log in with administrator account admin001.")
	}
}

// Close method helps you close BoltDB correctly.
func (myBolt *MyBolt) Close() {
	err := myBolt.db.Close()
	if err != nil {
		log.Printf("Failed to close BoltDB.")
	}
}

// Insert inserts data into a BoltDB bucket based on the provided BoltTable.
func (myBolt *MyBolt) Insert(tb core.TableData) error {
	var tableName string
	var key string
	var value []byte

	switch t := tb.(type) {
	case core.Account:
		tableName = myBolt.AccountTable
		key = t.UserId
		value, _ = json.Marshal(t)
	case core.UserGroup:
		tableName = myBolt.UserGroupTable
		key = t.GroupId
		value, _ = json.Marshal(t)
	case core.Role:
		tableName = myBolt.RoleTable
		key = t.RoleId
		value, _ = json.Marshal(t)
	default:
		return fmt.Errorf("unknown table type")
	}
	return myBolt.db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(tableName))
		if err != nil {
			return err
		}
		err = bucket.Put([]byte(key), value)
		if err != nil {
			return err
		}
		fmt.Printf("Inserted: tablename：%s, key: %s\n", tableName, key)
		return nil
	})
}

// Search 从 BoltDB 中检索数据并返回一个接口，该接口表示一个指向某个结构体指针的值。
// 具体的结构体类型取决于传递给函数的 table 参数的类型。
//
// 如果 table 参数是 core.Account 类型，返回值将是 *core.Account。
// 如果 table 参数是 core.UserGroup 类型，返回值将是 *core.UserGroup。
// 如果 table 参数是 core.Role 类型，返回值将是 *core.Role。
//
// 如果在检索或解析数据时发生错误，将返回错误。
func (myBolt *MyBolt) Search(tb core.TableData) (interface{}, error) {
	var tableName string
	var simpleStruct interface{}
	var key string
	switch t := tb.(type) {
	case *core.Account:
		simpleStruct = &core.Account{}
		tableName = myBolt.AccountTable
		key = t.UserId
	case core.Account:
		simpleStruct = &core.Account{}
		tableName = myBolt.AccountTable
		key = t.UserId
	case *core.UserGroup:
		simpleStruct = &core.UserGroup{}
		tableName = myBolt.UserGroupTable
		key = t.GroupId
	case core.UserGroup:
		simpleStruct = &core.UserGroup{}
		tableName = myBolt.UserGroupTable
		key = t.GroupId
	case *core.Role:
		simpleStruct = &core.Role{}
		tableName = myBolt.RoleTable
		key = t.RoleId
	case core.Role:
		simpleStruct = &core.Role{}
		tableName = myBolt.RoleTable
		key = t.RoleId
	default:
		return nil, fmt.Errorf("unknown table type")
	}
	err := myBolt.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(tableName))
		if bucket == nil {
			return nil // 桶不存在，返回空结果
		}
		data := bucket.Get([]byte(key))
		// 根据表的类型反序列化数据为相应的结构体
		if err := json.Unmarshal(data, simpleStruct); err != nil {
			return err
		}
		return nil
	})
	return simpleStruct, err
}

// CreateTable creates a new BoltDB bucket with the specified name within the database.
func (myBolt *MyBolt) CreateTable(tableName string) error {
	return myBolt.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte(tableName))
		return err
	})
}

// PrintAll shows all the data that has been loaded into the BoltDB.
func (myBolt *MyBolt) PrintAll(tableName string) error {
	var simpleStruct interface{}
	var err error

	switch tableName {
	case myBolt.AccountTable:
		simpleStruct = &core.Account{}
	case myBolt.UserGroupTable:
		simpleStruct = &core.UserGroup{}
	case myBolt.RoleTable:
		simpleStruct = &core.Role{}
	default:
		return fmt.Errorf("unknown bucket name")
	}
	// Create a table to display the data
	table := tablewriter.NewWriter(os.Stdout)
	// Start a BoltDB view transaction to read data
	err = myBolt.db.View(func(tx *bolt.Tx) error {
		// Retrieve the specified bucket
		bucket := tx.Bucket([]byte(tableName))
		if bucket == nil {
			return fmt.Errorf("Bucket '%service' not found.", tableName)
		}
		// Add struct field names as the first row (header)
		structFields := reflect.TypeOf(simpleStruct).Elem()
		header := make([]string, structFields.NumField())
		for i := 0; i < structFields.NumField(); i++ {
			header[i] = structFields.Field(i).Name
		}
		table.SetHeader(header)
		// Iterate through all key-value pairs in the bucket
		return bucket.ForEach(func(key, value []byte) error {
			// Initialize a new instance of the corresponding struct type
			newObj := reflect.New(reflect.TypeOf(simpleStruct).Elem()).Interface()
			// Unmarshal the value into the struct
			if err := json.Unmarshal(value, newObj); err != nil {
				return err
			}
			// Get struct values and add them as a row in the table
			structValues := reflect.ValueOf(newObj).Elem()
			row := make([]string, structValues.NumField())
			for i := 0; i < structValues.NumField(); i++ {
				row[i] = fmt.Sprintf("%v", structValues.Field(i).Interface())
			}
			table.Append(row)
			return nil
		})
	})
	if err != nil {
		return err
	}
	// Render the table
	table.Render()
	return nil
}

// IsExist Determines whether a key exists
func (myBolt *MyBolt) IsExist(tableName string, id string) (bool, error) {
	var exist bool
	switch tableName {
	case myBolt.AccountTable, myBolt.UserGroupTable, myBolt.RoleTable:
		err := myBolt.db.View(func(tx *bolt.Tx) error {
			bucket := tx.Bucket([]byte(tableName)) // 替换成你的桶名
			key := []byte(id)                      // 替换成你要查询的键名
			val := bucket.Get(key)
			if val == nil {
				exist = false
			} else {
				exist = true
			}
			return nil
		})
		if err != nil {
			return false, err
		} else {
			return exist, nil
		}
	default:
		return false, fmt.Errorf("unknown bucket name")
	}
}

// Update Helps users edit changeable properties
func (myBolt *MyBolt) Update(tb core.TableData) error {
	var bucketName string
	var key string
	var value []byte
	switch t := tb.(type) {
	case *core.Account:
		bucketName = myBolt.AccountTable
		key = t.UserId
		value, _ = json.Marshal(t)
	case core.Account:
		bucketName = myBolt.AccountTable
		key = t.UserId
		value, _ = json.Marshal(t)
	case *core.UserGroup:
		bucketName = myBolt.UserGroupTable
		key = t.GroupId
		value, _ = json.Marshal(t)
	case core.UserGroup:
		bucketName = myBolt.UserGroupTable
		key = t.GroupId
		value, _ = json.Marshal(t)
	case *core.Role:
		bucketName = myBolt.RoleTable
		key = t.RoleId
		value, _ = json.Marshal(t)
	case core.Role:
		bucketName = myBolt.RoleTable
		key = t.RoleId
		value, _ = json.Marshal(t)
	default:
		return fmt.Errorf("unknown table type")
	}
	return myBolt.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return fmt.Errorf("bucket not found")
		}
		existingValue := bucket.Get([]byte(key))
		if existingValue == nil {
			return fmt.Errorf("record not found")
		}
		// update
		return bucket.Put([]byte(key), value)
	})
}

// Del Helps administrator delete a common user
func (myBolt *MyBolt) Del(tb core.TableData) error {
	var bucketName string
	var key string
	switch t := tb.(type) {
	case *core.Account:
		bucketName = myBolt.AccountTable
		key = t.UserId
	case core.Account:
		bucketName = myBolt.AccountTable
		key = t.UserId
	case *core.UserGroup:
		bucketName = myBolt.UserGroupTable
		key = t.GroupId
	case core.UserGroup:
		bucketName = myBolt.UserGroupTable
		key = t.GroupId
	case *core.Role:
		bucketName = myBolt.RoleTable
		key = t.RoleId
	case core.Role:
		bucketName = myBolt.RoleTable
		key = t.RoleId
	default:
		return fmt.Errorf("unknown table type")
	}
	return myBolt.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return fmt.Errorf("bucket not found")
		}
		// delete
		return bucket.Delete([]byte(key))
	})
}
