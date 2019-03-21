package dao

import "container/list"

type DAOFilterInterface interface {
	ToSql() string
}

type DAOInterface interface {
	Create() int
	find(id int) *list.List
	Update() bool
	Delete(id int) bool
	List(filter DAOFilterInterface, offset int, limit int) *list.List
}
