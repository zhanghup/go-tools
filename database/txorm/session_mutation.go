package txorm

import (
	"strings"
)

func (this *Session) Insert(bean ...any) error {
	return this.AutoClose(func() error {
		this.Table(bean)
		_, err := this.sess.Insert(bean...)
		return err
	})
}

func (this *Session) Update(bean any, condiBean ...any) error {
	return this.AutoClose(func() error {
		this.Table(bean)
		sqlstr := strings.TrimSpace(this._sql(false))
		_, err := this.sess.Table(this.tableName).Where(sqlstr, this.args...).Update(bean, condiBean...)
		return err
	})
}

func (this *Session) Delete(bean ...any) error {
	return this.AutoClose(func() error {
		this.Table(bean)
		sqlstr := strings.TrimSpace(this._sql(false))

		_, err := this.sess.Table(this.tableName).Where(sqlstr, this.args...).Delete(bean...)
		return err
	})
}

func (this *Session) Exec() error {
	return this.AutoClose(func() error {
		sqls := []any{this._sql_with() + " " + this._sql(true)}
		_, err := this.sess.Exec(append(sqls, this.args...)...)
		return err
	})

}
