package txorm

import (
	"fmt"
	"strings"
)

func (this *Session) SelectSql(bean interface{}, orderFlag bool, columns ...string) string {
	this.Table(bean)
	sqlstr := strings.TrimSpace(this._sql(orderFlag))
	if strings.Index(sqlstr, "select") == 0 || strings.Index(sqlstr, "SELECT") == 0 {
		if len(columns) > 0 {
			column := strings.Join(columns, ",")
			return fmt.Sprintf("%s select %s from ( %s) _", this._sql_with(), column, sqlstr)
		}

		return fmt.Sprintf("%s %s", this._sql_with(), sqlstr)
	}

	column := "*"
	if len(columns) > 0 {
		column = strings.Join(columns, ",")
	}

	switch {
	case strings.Index(sqlstr, "limit") == 0,
		strings.Index(sqlstr, "where") == 0,
		strings.Index(sqlstr, "group") == 0,
		strings.Index(sqlstr, "order") == 0,
		sqlstr == "":
		return fmt.Sprintf("%s select %s from %s %s", this._sql_with(), column, this.tableName, sqlstr)
	default:
		return fmt.Sprintf("%s select %s from %s where %s", this._sql_with(), column, this.tableName, sqlstr)
	}

}

func (this *Session) Find(bean interface{}) error {
	return this.AutoClose(func() error {
		return this.sess.SQL(this.SelectSql(bean, true), this.args...).Find(bean)
	})
}

func (this *Session) Get(bean interface{}) (v bool, err error) {
	err = this.AutoClose(func() error {
		v, err = this.sess.SQL(this.SelectSql(bean, true)+" limit 1", this.args...).Get(bean)
		return err
	})
	return
}

//func (this *Session) Map() (v []map[string]interface{}, err error) {
//	err = this.AutoClose(func() error {
//		rows, err := this.sess.DB().Query(this.SelectSql(nil, true), this.args...)
//		if err != nil {
//			return err
//		}
//
//		for rows.Next() {
//			vv := map[string]interface{}{}
//			if err = rows.ScanMap(&vv); err != nil {
//				return err
//			} else {
//				v = append(v, vv)
//			}
//		}
//		return nil
//	})
//	return
//}

func (this *Session) Map() (v []map[string]interface{}, err error) {
	err = this.AutoClose(func() error {
		v, err = this.sess.SQL(this.SelectSql(nil, true), this.args...).QueryInterface()
		return err
	})
	return
}

func (this *Session) Exists() (v bool, err error) {
	err = this.AutoClose(func() error {
		v, err = this.sess.SQL(this.SelectSql(nil, false, "1")+" limit 1", this.args...).Exist()
		return err
	})
	return
}
