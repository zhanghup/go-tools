package txorm

import (
	"context"
	"strings"
)
func (this *Session) Insert(bean ...interface{}) error {
	return this.begin(func() error {
		this.Table(bean)
		_, err := this.sess.Insert(bean...)
		return err
	})
}

func (this *Session) Update(bean interface{}, condiBean ...interface{}) error {
	return this.begin(func() error {
		this.Table(bean)
		sqlstr := strings.TrimSpace(this._sql(false))
		_, err := this.sess.Table(this.tableName).Where(sqlstr, this.args...).Update(bean, condiBean...)
		return err
	})
}

func (this *Session) Delete(bean ...interface{}) error {
	return this.begin(func() error {
		this.Table(bean)
		sqlstr := strings.TrimSpace(this._sql(false))

		_, err := this.sess.Table(this.tableName).Where(sqlstr, this.args...).Delete(bean...)
		return err
	})
}

func (this *Session) TS(fn func(ctx context.Context, sess ISession) error) error {
	return this.AutoClose(func() error {
		this.Begin()
		err := fn(this.context, this)
		if err != nil {
			_ = this.Rollback()
			return err
		} else {
			err = this.Commit()
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (this *Session) Exec() error {
	return this.begin(func() error {
		sqls := []interface{}{this._sql_with() + " " + this._sql(true)}
		_, err := this.sess.Exec(append(sqls, this.args...)...)
		return err
	})

}
