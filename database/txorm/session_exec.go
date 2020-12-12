package txorm

import (
	"fmt"
	"github.com/zhanghup/go-tools"
)

func (this *Session) Insert(bean ...interface{}) error {
	if this.autoClose {
		// 由engine直接进入的方法，需要自动关闭session
		defer this.AutoClose()
	}
	_, err := this.sess.Insert(bean...)
	return err
}

func (this *Session) Update(bean interface{}, condiBean ...interface{}) error {
	if this.autoClose {
		// 由engine直接进入的方法，需要自动关闭session
		defer this.AutoClose()
	}
	_, err := this.sess.Update(bean, condiBean...)
	return err
}

func (this *Session) Delete(bean interface{}) error {
	if this.autoClose {
		// 由engine直接进入的方法，需要自动关闭session
		defer this.AutoClose()
	}
	_, err := this.sess.Delete(bean)
	return err
}

func (this *Session) TS(fn func(sess ISession) error, commit ...bool) error {
	err := this.Begin()
	if err != nil {
		return err
	}
	err = fn(this)
	if err != nil {
		_ = this.Rollback()
		return err
	}
	if len(commit) > 0 && commit[0] {
		return this.Commit()
	} else if this.autoClose {
		return this.AutoClose()
	}
	return nil
}

func (this *Session) Find(bean interface{}) error {
	if this.autoClose {
		// 由engine直接进入的方法，需要自动关闭session
		defer this.AutoClose()
	}
	return this.sess.SQL(this.sqlwith+" "+this.sql, this.args...).Find(bean)
}

func (this *Session) Get(bean interface{}) (bool, error) {
	if this.autoClose {
		// 由engine直接进入的方法，需要自动关闭session
		defer this.AutoClose()
	}
	return this.sess.SQL(this.sqlwith+" "+this.sql, this.args...).Get(bean)

}

func (this *Session) Page(index, size int, count bool, bean interface{}) (int, error) {
	if this.autoClose {
		// 由engine直接进入的方法，需要自动关闭session
		defer this.AutoClose()
	}
	if size < 0 {
		err := this.sess.SQL(this.sql, this.args...).Find(bean)
		return 0, err
	} else if size == 0 {
		total := 0
		_, err := this.sess.SQL(fmt.Sprintf("%s select count(1) from (%s) _", this.sqlwith, this.sql), this.args...).Get(&total)
		return total, err
	} else {
		err := this.sess.SQL(fmt.Sprintf("%s limit ?,?", this.sql), append(this.args, (index-1)*size, size)...).Find(bean)
		if err != nil {
			return 0, err
		}
		if count {
			total := 0
			_, err := this.sess.SQL(fmt.Sprintf("%s select count(1) from (%s) _", this.sqlwith, this.sql), this.args...).Get(&total)
			return total, err
		}
	}

	return 0, nil
}

func (this *Session) Page2(index, size *int, count *bool, bean interface{}) (int, error) {
	if index == nil {
		index = tools.Ptr.Int(1)
	}
	if size == nil {
		size = tools.Ptr.Int(1)
	}
	if count == nil {
		count = tools.Ptr.Bool(false)
	}
	return this.Page(*index, *size, *count, bean)
}

func (this *Session) Exec() error {
	if this.autoClose {
		// 由engine直接进入的方法，需要自动关闭session
		defer this.AutoClose()
	}

	sqls := []interface{}{this.sql}
	_, err := this.sess.Exec(append(sqls, this.args...)...)
	return err
}
