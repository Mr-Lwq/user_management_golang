package dao

import "user_management_golang/src"

type ORM interface {
	Close()
	Insert(tb src.TableData) error
	Search(tb src.TableData) (interface{}, error)
	Update(tb src.TableData) error
	Del(tb src.TableData) error
}
