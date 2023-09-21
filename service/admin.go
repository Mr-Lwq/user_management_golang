package service

import (
	"fmt"
	"unicode/utf8"
	"user_management_golang/dao"
	"user_management_golang/src"
)

// register
func UserRegister(account src.Account) (bool, error) {
	if mybolt, err := dao.NewMyBolt(); err != nil {
		// 处理错误，例如打印错误信息或返回错误响应
		fmt.Printf("Failed to create MyBolt: %v\n", err)
	} else {
		defer mybolt.Close()
		exist, err := mybolt.IsExistQuery(mybolt.AccountTable, account.UserId)
		if err != nil {
			return false, err
		} else {
			if exist {
				return false, fmt.Errorf("registration failed because the user name already exists")
			} else {
				passwordLength := utf8.RuneCountInString(account.Password)
				usernameLength := utf8.RuneCountInString(account.Username)
				if usernameLength >= 8 && usernameLength <= 16 && passwordLength >= 8 && passwordLength <= 16 {
					if err = mybolt.InsertData(account); err != nil {
						return false, err
					} else {
						return true, nil
					}
				} else {
					return false, fmt.Errorf("username and password must be a string of 8 to 16 characters")
				}
			}
		}
	}
	return false, nil
}

type Admin struct {
	src.Account
}
