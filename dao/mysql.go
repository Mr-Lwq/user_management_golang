/*
Package dao is the Dao layer of the program.
This program uses two databases, BoltDB and MySQL.
If you want to keep it simple, starting Simple Mode will default to the embedded
database BoltDB.
*/

package dao

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"
	"user_management_golang/src"
)

type MysqlCfg struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

type Mysql struct {
	db             *sql.DB
	UserGroupTable string
	RoleTable      string
	AccountTable   string
}

// NewMysql create a new Mysql instance for manage mysql database.
func NewMysql(cfg MysqlCfg) (*Mysql, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	// 测试连接是否成功
	if err := db.Ping(); err != nil {
		return nil, err
	}
	// 返回 Mysql 结构体实例
	mysql := &Mysql{
		db:             db,
		UserGroupTable: "user_group_table",
		RoleTable:      "role_table",
		AccountTable:   "account_table",
	}
	mysql.init()
	return mysql, nil
}

// init helps you initialize the Mysql DB table the first time you open the program.
func (mysql *Mysql) init() {
	tables := []string{
		mysql.UserGroupTable,
		mysql.RoleTable,
		mysql.AccountTable,
	}
	// Check whether you need to create a table
	firstMake := false
	for _, table := range tables {
		query := fmt.Sprintf("SELECT 1 FROM %s LIMIT 1", table)
		_, err := mysql.db.Exec(query)
		if err != nil {
			firstMake = true
			break
		}
	}
	if firstMake {
		for _, table := range tables {
			err := mysql.CreateTable(table)
			if err != nil {
				log.Printf("Failed to create table %s: %v", table, err)
				return
			}
		}
		fmt.Println("First setting beginning...")
		admin := src.Account{
			UserId:         "admin",
			Username:       "admin",
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
			log.Fatal(err)
		}
		fmt.Println("[UserId: admin, Password: 88888888] Create account table success. Add root account.")
		userGroup := src.UserGroup{
			GroupId:       "administrators",
			GroupName:     "administrators",
			GroupLeads:    "admin",
			Description:   "",
			Permissions:   []string{},
			Members:       []string{"admin"},
			CreatedAt:     time.Now(),
			LastUpdatedAt: time.Now(),
		}
		err = mysql.Insert(userGroup)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("[GroupId: group, GroupName: administrators] Create user_group table success. Add administrator group.")
		role := src.Role{
			RoleId:      "admin",
			RoleName:    "admin",
			Description: "",
			Permissions: []string{},
			CreatedAt:   time.Now(),
		}
		err = mysql.Insert(role)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("[RoleId: role, RoleName: admin] Create role table success. Add admin role.")
		fmt.Println("The first setting is successful. Please log in with administrator account admin.")
	}
}

// Close helps you close MysqlDB correctly.
func (mysql *Mysql) Close() {
	if mysql.db != nil {
		err := mysql.db.Close()
		if err != nil {
			log.Printf("Failed to close MySql.")
		}
	}
}

// Insert helps you insert a new record into mysql db.
func (mysql *Mysql) Insert(tb src.TableData) error {
	tableType := tb.GetTableType()
	var tableName string
	switch tableType {
	case "Account":
		tableName = mysql.AccountTable
	case "UserGroup":
		tableName = mysql.UserGroupTable
	case "Role":
		tableName = mysql.RoleTable
	default:
		return fmt.Errorf("unknown table type")
	}
	dataMap, err := structToMap(tb)
	if err != nil {
		return err
	}
	columns := make([]string, 0)
	values := make([]string, 0)
	args := make([]interface{}, 0)

	for column, value := range dataMap {
		switch v := value.(type) {
		case []string:
			// 如果字段类型是 []string，将其拼接成由逗号连接的字符串
			strValue := strings.Join(v, ",")
			columns = append(columns, column)
			values = append(values, "?")
			args = append(args, strValue)
		default:
			columns = append(columns, column)
			values = append(values, "?")
			args = append(args, value)
		}
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, strings.Join(columns, ","), strings.Join(values, ","))

	result, err := mysql.db.Exec(query, args...)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Printf("Inserted %d rows into table %s\n", rowsAffected, tableName)
	return nil
}

// Update helps you update existing records from the mysql db.
func (mysql *Mysql) Update(tb src.TableData) error {
	var tableName string
	var idValue string
	var idColName string
	switch t := tb.(type) {
	case src.Account:
		tableName = mysql.AccountTable
		idValue = t.UserId
		idColName = "UserId"
	case src.UserGroup:
		tableName = mysql.UserGroupTable
		idValue = t.GroupId
		idColName = "GroupId"
	case src.Role:
		tableName = mysql.RoleTable
		idValue = t.RoleId
		idColName = "RoleId"
	default:
		return fmt.Errorf("unknown table type")
	}
	dataMap, err := structToMap(tb)
	if err != nil {
		return err
	}
	// Define a slice to hold the update expressions for each column
	updateExpressions := make([]string, 0)
	args := make([]interface{}, 0)

	for column, value := range dataMap {
		// Skip the primary key column if it's empty (assuming your table has one)
		if column == idColName && value == nil {
			continue
		}
		switch v := value.(type) {
		case []string:
			// If the field type is []string, join it into a comma-separated string
			strValue := strings.Join(v, ",")
			updateExpressions = append(updateExpressions, fmt.Sprintf("%s = ?", column))
			args = append(args, strValue)
		default:
			updateExpressions = append(updateExpressions, fmt.Sprintf("%s = ?", column))
			args = append(args, value)
		}
	}

	if len(updateExpressions) == 0 {
		return fmt.Errorf("no columns to update")
	}

	// Construct the UPDATE query
	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s = ?", tableName, strings.Join(updateExpressions, ", "), idColName)

	args = append(args, idValue)
	result, err := mysql.db.Exec(query, args...)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Printf("Updated %d rows in table %s\n", rowsAffected, tableName)
	return nil
}

// Del helps you delete existing records from the mysql db.
func (mysql *Mysql) Del(tb src.TableData) error {
	var tableName string
	var idValue string
	var idColName string
	switch t := tb.(type) {
	case src.Account:
		tableName = mysql.AccountTable
		idValue = t.UserId
		idColName = "UserId"
	case src.UserGroup:
		tableName = mysql.UserGroupTable
		idValue = t.GroupId
		idColName = "GroupId"
	case src.Role:
		tableName = mysql.RoleTable
		idValue = t.RoleId
		idColName = "RoleId"
	default:
		return fmt.Errorf("unknown table type")
	}
	// 构建 DELETE 查询
	query := fmt.Sprintf("DELETE FROM %s WHERE %s = ?", tableName, idColName)

	// 执行 DELETE 查询
	_, err := mysql.db.Exec(query, idValue)
	if err != nil {
		return err
	}

	fmt.Printf("Deleted row from table %s where %s = %v\n", tableName, idColName, idValue)
	return nil
}

// Search helps you query the existing records in the database.
func (mysql *Mysql) Search(tb src.TableData) (interface{}, error) {
	var tableName string
	var idValue string
	var idColName string
	var simpleStruct interface{}
	var err error
	switch t := tb.(type) {
	case src.Account:
		tableName = mysql.AccountTable
		idValue = t.UserId
		idColName = "UserId"
	case src.UserGroup:
		tableName = mysql.UserGroupTable
		idValue = t.GroupId
		idColName = "GroupId"
	case src.Role:
		tableName = mysql.RoleTable
		idValue = t.RoleId
		idColName = "RoleId"
	default:
		return nil, fmt.Errorf("unknown table type")
	}

	// Construct the query
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s = ?", tableName, idColName)

	// Execute the query
	rows, err := mysql.db.Query(query, idValue)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	// Check if there is at least one row in the result
	if rows.Next() {
		switch tableName {
		case mysql.AccountTable:
			var acount src.Account
			var rolesStr, groupsStr, lastLoginAtStr, createAtStr, updatedAtStr string
			err = rows.Scan(
				&acount.Password,
				&lastLoginAtStr,
				&groupsStr,
				&acount.FullName,
				&acount.SessionToken,
				&acount.ProfilePicture,
				&acount.UserId,
				&acount.Phone,
				&updatedAtStr,
				&acount.Username,
				&acount.Email,
				&rolesStr,
				&acount.Status,
				&createAtStr,
			)
			if err != nil {
				return nil, err
			}
			updatedAt, _ := time.Parse("2006-01-02 15:04:05", updatedAtStr)
			createAt, _ := time.Parse("2006-01-02 15:04:05", createAtStr)
			lastLoginAt, _ := time.Parse("2006-01-02 15:04:05", lastLoginAtStr)
			acount.CreatedAt = createAt
			acount.LastLoginAt = lastLoginAt
			acount.UpdatedAt = updatedAt
			acount.Roles = strings.Split(rolesStr, ",")
			acount.UserGroups = strings.Split(groupsStr, ",")
			simpleStruct = acount
			return simpleStruct, nil

		case mysql.UserGroupTable:
			var group src.UserGroup
			var lastUpdatedAtStr, createAtStr, permissionStr, membersStr string
			err = rows.Scan(
				&createAtStr,
				&lastUpdatedAtStr,
				&group.GroupId,
				&group.GroupName,
				&group.Description,
				&permissionStr,
				&membersStr,
			)
			if err != nil {
				return nil, err
			}
			lastUpdatedAt, _ := time.Parse("2006-01-02 15:04:05", lastUpdatedAtStr)
			createAt, _ := time.Parse("2006-01-02 15:04:05", createAtStr)
			permission := strings.Split(permissionStr, ",")
			members := strings.Split(membersStr, ",")
			group.LastUpdatedAt = lastUpdatedAt
			group.CreatedAt = createAt
			group.Permissions = permission
			group.Members = members
			simpleStruct = group
			return simpleStruct, nil

		case mysql.RoleTable:
			var role src.Role
			var createAtStr, permissionStr string

			err = rows.Scan(
				&role.Description,
				&permissionStr,
				&createAtStr,
				&role.RoleId,
				&role.RoleName,
			)
			if err != nil {
				return nil, err
			}
			createAt, _ := time.Parse("2006-01-02 15:04:05", createAtStr)
			permission := strings.Split(permissionStr, ",")
			role.CreatedAt = createAt
			role.Permissions = permission
			simpleStruct = role
			return simpleStruct, nil

		default:
			return nil, fmt.Errorf("no record found for the given primary key")
		}
	} else {
		return nil, fmt.Errorf("no record found for the given primary key")
	}
}

func (mysql *Mysql) CreateTable(tableName string) error {
	var tableStructure map[string]string
	var primaryKey string
	var indexes []string

	switch tableName {
	case mysql.AccountTable:
		tableStructure = map[string]string{
			"UserId":         "VARCHAR(255)",
			"Username":       "VARCHAR(255)",
			"Password":       "VARCHAR(255)",
			"Email":          "VARCHAR(255)",
			"Phone":          "VARCHAR(255)",
			"FullName":       "VARCHAR(255)",
			"Roles":          "VARCHAR(255)",
			"Status":         "VARCHAR(255)",
			"CreatedAt":      "TIMESTAMP",
			"UpdatedAt":      "TIMESTAMP",
			"LastLoginAt":    "TIMESTAMP",
			"SessionToken":   "VARCHAR(255)",
			"ProfilePicture": "VARCHAR(255)",
			"UserGroups":     "VARCHAR(255)",
		}
		primaryKey = "UserId"
		indexes = []string{"Username", "Email"}
	case mysql.UserGroupTable:
		tableStructure = map[string]string{
			"GroupId":       "VARCHAR(255)",
			"GroupLeads":    "VARCHAR(255)",
			"GroupName":     "VARCHAR(255)",
			"Description":   "VARCHAR(255)",
			"Permissions":   "VARCHAR(255)",
			"Members":       "VARCHAR(255)",
			"CreatedAt":     "TIMESTAMP",
			"LastUpdatedAt": "TIMESTAMP",
		}
		primaryKey = "GroupId"
		indexes = []string{"GroupName"}
	case mysql.RoleTable:
		tableStructure = map[string]string{
			"RoleId":      "VARCHAR(255)",
			"RoleName":    "VARCHAR(255)",
			"Description": "VARCHAR(255)",
			"Permissions": "VARCHAR(255)",
			"CreatedAt":   "TIMESTAMP",
		}
		primaryKey = "RoleId"
		indexes = []string{"RoleName"}
	default:
		return fmt.Errorf("unknown table name")
	}

	createTableQuery := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (", tableName)
	for columnName, columnType := range tableStructure {
		createTableQuery += fmt.Sprintf("%s %s,", columnName, columnType)
	}
	createTableQuery += fmt.Sprintf("PRIMARY KEY (%s),", primaryKey)

	for _, index := range indexes {
		createTableQuery += fmt.Sprintf("INDEX %s_index (%s),", index, index)
	}

	createTableQuery = createTableQuery[:len(createTableQuery)-1] + ")"

	// 执行 CREATE TABLE 语句
	_, err := mysql.db.Exec(createTableQuery)
	if err != nil {
		return err
	}
	return nil
}

func structToMap(data interface{}) (map[string]interface{}, error) {
	dataMap := make(map[string]interface{})

	dataType := reflect.TypeOf(data)
	dataValue := reflect.ValueOf(data)

	if dataType.Kind() == reflect.Ptr {
		dataType = dataType.Elem()
		dataValue = dataValue.Elem()
	}

	if dataType.Kind() != reflect.Struct {
		return nil, fmt.Errorf("data is not a struct")
	}

	for i := 0; i < dataType.NumField(); i++ {
		field := dataType.Field(i)
		value := dataValue.Field(i)
		if field.PkgPath != "" {
			continue
		}
		fieldName := getFieldTag(field)
		dataMap[fieldName] = value.Interface()
	}
	return dataMap, nil
}

func getFieldTag(field reflect.StructField) string {
	tag := field.Tag.Get("json")
	if tag == "" {
		return field.Name
	}
	parts := strings.Split(tag, ",")
	return parts[0]
}
