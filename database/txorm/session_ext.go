package txorm

import (
	"fmt"
	"github.com/zhanghup/go-tools"
)

func (this *Session) Count() (total int64, err error) {
	err = this.AutoClose(func() error {
		total, err = this.sess.SQL(this._sql_with()+" "+this._sql(), this.args...).Count()
		return err
	})
	return

}

func (this *Session) Int() (n int, err error) {
	err = this.AutoClose(func() error {
		_, err = this.Get(&n)
		return err
	})
	return
}

func (this *Session) Int64() (n int64, err error) {
	err = this.AutoClose(func() error {
		_, err = this.Get(&n)
		return err
	})
	return
}

func (this *Session) Float64() (f float64, err error) {
	err = this.AutoClose(func() error {
		_, err = this.Get(&f)
		return err
	})

	return
}

func (this *Session) String() (v string, err error) {
	err = this.AutoClose(func() error {
		_, err = this.Get(&v)
		return err
	})
	return
}

func (this *Session) Strings() (v []string, err error) {
	err = this.AutoClose(func() error {
		err = this.Find(&v)
		return err
	})
	return

}

// Page 分页查询
// size < 0 查询所有
// size = 0 只查询所有数据的量，不查询具体数据
// count = true 分页查询数据并且查询数据总量
func (this *Session) Page(index, size int, count bool, bean interface{}) (v int, err error) {
	err = this.AutoClose(func() error {
		if size < 0 {
			err = this.sess.SQL(this._sql_with()+" "+this._sql(), this.args...).Find(bean)
			return err
		} else if size == 0 {
			_, err = this.sess.SQL(fmt.Sprintf("%s select count(1) from (%s) _", this._sql_with(), this.sql), this.args...).Get(&v)
			return err
		} else {
			err = this.sess.SQL(fmt.Sprintf("%s limit ?,?", this._sql_with()+" "+this._sql()), append(this.args, (index-1)*size, size)...).Find(bean)
			if err != nil {
				return err
			}
			if count {
				_, err = this.sess.SQL(fmt.Sprintf("%s select count(1) from (%s) _", this._sql_with(), this.sql), this.args...).Get(&v)
				return err
			}
		}
		return nil
	})
	return
}

func (this *Session) Page2(index, size *int, count *bool, bean interface{}) (int, error) {
	if index == nil {
		index = tools.PtrOfInt(1)
	}
	if size == nil {
		size = tools.PtrOfInt(1)
	}
	if count == nil {
		count = tools.PtrOfBool(false)
	}
	return this.Page(*index, *size, *count, bean)
}
