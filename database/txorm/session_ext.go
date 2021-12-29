package txorm

import (
	"fmt"
	"github.com/zhanghup/go-tools"
)

func (this *Session) Count() (int64, error) {
	if this.autoClose {
		// 由engine直接进入的方法，需要自动关闭session
		defer this.AutoClose()
	}

	return this.sess.SQL(this._sql_with()+" "+this._sql(), this.args...).Count()
}

func (this *Session) Int() (int, error) {
	if this.autoClose {
		// 由engine直接进入的方法，需要自动关闭session
		defer this.AutoClose()
	}

	n := 0
	_, err := this.Get(&n)
	return n, err
}

func (this *Session) Int64() (int64, error) {
	if this.autoClose {
		// 由engine直接进入的方法，需要自动关闭session
		defer this.AutoClose()
	}

	n := int64(0)
	_, err := this.Get(&n)
	return n, err
}

func (this *Session) Float64() (float64, error) {
	if this.autoClose {
		// 由engine直接进入的方法，需要自动关闭session
		defer this.AutoClose()
	}

	n := float64(0)
	_, err := this.Get(&n)
	return n, err
}

func (this *Session) String() (string, error) {
	if this.autoClose {
		// 由engine直接进入的方法，需要自动关闭session
		defer this.AutoClose()
	}

	n := ""
	_, err := this.Get(&n)
	return n, err
}

func (this *Session) Strings() ([]string, error) {
	if this.autoClose {
		// 由engine直接进入的方法，需要自动关闭session
		defer this.AutoClose()
	}

	n := make([]string, 0)
	err := this.Find(&n)
	return n, err
}

func (this *Session) Page(index, size int, count bool, bean interface{}) (int, error) {
	if this.autoClose {
		// 由engine直接进入的方法，需要自动关闭session
		defer this.AutoClose()
	}
	if size < 0 {
		err := this.sess.SQL(this._sql_with()+" "+this._sql(), this.args...).Find(bean)
		return 0, err
	} else if size == 0 {
		total := 0
		_, err := this.sess.SQL(fmt.Sprintf("%s select count(1) from (%s) _", this._sql_with(), this.sql), this.args...).Get(&total)
		return total, err
	} else {
		err := this.sess.SQL(fmt.Sprintf("%s limit ?,?", this._sql_with()+" "+this._sql()), append(this.args, (index-1)*size, size)...).Find(bean)
		if err != nil {
			return 0, err
		}
		if count {
			total := 0
			_, err := this.sess.SQL(fmt.Sprintf("%s select count(1) from (%s) _", this._sql_with(), this.sql), this.args...).Get(&total)
			return total, err
		}
	}

	return 0, nil
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
