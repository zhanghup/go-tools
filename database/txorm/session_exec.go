package txorm

import (
	"fmt"
	"github.com/zhanghup/go-tools"
	"regexp"
	"strings"
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

func (this *Session) Order(order ...string) ISession {
	this.orderby = order
	return this
}

func (this *Session) _sql() string {
	if len(this.orderby) == 0 {
		return this.sql
	}
	res := regexp.MustCompile(`\(.*\)`).ReplaceAllString(this.sql, "")
	match := regexp.MustCompile(`order\s+by\s+`).MatchString(res)

	orderBy := make([]string, 0)
	for _, s := range this.orderby {
		if regexp.MustCompile(`^-[a-zA-Z0-9_]+`).MatchString(s) {
			ss := strings.Replace(s, "-", "", 1)
			orderBy = append(orderBy, ss+" desc")
		} else if regexp.MustCompile(`[a-zA-Z0-9_]+`).MatchString(s) {
			orderBy = append(orderBy, s+" asc")
		} else {
			orderBy = append(orderBy, s+" ")
		}
	}
	if match {
		return this.sql + "," + strings.Join(orderBy, ",")
	} else {
		return this.sql + " order by " + strings.Join(orderBy, ",")
	}
}

func (this *Session) _sql_with() string {
	sqlwith := ""
	if len(this.withs) > 0 {
		// 去重
		withs := []string{"\n with recursive _ as (select 1)"}
		wmap := map[string]bool{}
		for _, w := range this.withs {
			wmap[w] = true
		}
		for k := range wmap {
			kk := tools.StrTmp(fmt.Sprintf("{{ sql_with_%s .ctx }}", k), map[string]interface{}{"ctx": this.Ctx()}).FuncMap(this.tmpWiths).String()
			withs = append(withs, fmt.Sprintf("%s as (%s)", k, kk))
		}

		sqlwith = strings.Join(withs, ",")
		sqlwith = tools.StrTmp(sqlwith, map[string]interface{}{"ctx": this.Ctx()}).FuncMap(this.tmpWiths).String()
	}
	return sqlwith
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
	return this.sess.SQL(this._sql_with()+" "+this._sql(), this.args...).Find(bean)
}

func (this *Session) Get(bean interface{}) (bool, error) {
	if this.autoClose {
		// 由engine直接进入的方法，需要自动关闭session
		defer this.AutoClose()
	}
	return this.sess.SQL(this._sql_with()+" "+this._sql(), this.args...).Get(bean)

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

func (this *Session) Map() ([]map[string][]byte, error) {
	if this.autoClose {
		// 由engine直接进入的方法，需要自动关闭session
		defer this.AutoClose()
	}

	return this.sess.SQL(this._sql_with()+" "+this._sql(), this.args...).Query()
}

func (this *Session) Exists() (bool, error) {
	if this.autoClose {
		// 由engine直接进入的方法，需要自动关闭session
		defer this.AutoClose()
	}

	return this.sess.SQL(this._sql_with()+" "+this._sql(), this.args...).Exist()
}

func (this *Session) Exec() error {
	if this.autoClose {
		// 由engine直接进入的方法，需要自动关闭session
		defer this.AutoClose()
	}

	sqls := []interface{}{this._sql_with() + " " + this._sql()}
	_, err := this.sess.Exec(append(sqls, this.args...)...)
	return err
}

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
