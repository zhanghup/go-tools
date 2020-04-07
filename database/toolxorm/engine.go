package toolxorm

import (
	"xorm.io/xorm"
)

type Engine struct {
	DB *xorm.Engine
}


func (this *Engine) NewSession() *Session {
	return &Session{Sess: this.DB.NewSession(), autoClose: false}
}

func (this *Engine) TS(fn func(sess *Session) error) {
	sess := &Session{Sess: this.DB.NewSession(), autoClose: false}
	sess.TS(fn)
}

func (this *Engine) SF(sql string, querys ...map[string]interface{}) *Session {
	sess := this.DB.NewSession()
	return (&Session{Sess: sess, autoClose: true}).SF(sql, querys...)
}

