package toolxorm

import "xorm.io/xorm"

type Session struct {
	Sess      *xorm.Session
	sql       string
	query     map[string]interface{}
	args      []interface{}
	autoClose bool
}
