package dao

import "user_management_golang/core"

type ORM interface {
	Close()
	Insert(tb core.TableData) error
	Search(tb core.TableData) (interface{}, error)
	Update(tb core.TableData) error
	Del(tb core.TableData) error
}
