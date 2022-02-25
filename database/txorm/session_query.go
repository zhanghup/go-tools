package txorm

import (
	"fmt"
	"github.com/zhanghup/go-tools"
	"strings"
	"time"
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

func (this *Session) Map() (v []map[string]interface{}, err error) {
	err = this.AutoClose(func() error {
		rows, err := this.sess.DB().Query(this.SelectSql(nil, true), this.args...)
		if err != nil {
			return err
		}

		types, err := rows.ColumnTypes()
		if err != nil {
			return err
		}

		for rows.Next() {
			vv := map[string][]byte{}
			if err = rows.ScanMap(&vv); err != nil {
				return err
			} else {
				vi := map[string]interface{}{}

				for _, o := range types {
					if colValue, ok := vv[o.Name()]; ok && colValue != nil {
						newValue := string(colValue)

						switch o.DatabaseTypeName() {
						case "DateTime", "DATETIME":
							vi[o.Name()], err = time.ParseInLocation("2006-01-02 15:04:05", newValue, time.Local)
							if err != nil {
								return err
							}
						case "Bool", "BOOL":
							vi[o.Name()] = colValue[0] == 0
						case "Blob", "BLOB":
							vi[o.Name()] = colValue
						case "Float", "FLOAT":
							vi[o.Name()] = tools.StrToFloat32(newValue)
						case "Double", "DOUBLE":
							vi[o.Name()] = tools.StrToFloat64(newValue)
						case "BigInt", "BIGINT":
							vi[o.Name()] = tools.StrToInt64(newValue)
						case "Int", "INT":
							vi[o.Name()] = tools.StrToInt(newValue)
						default:
							vi[o.Name()] = string(colValue)
						}
					} else {
						vi[o.Name()] = nil
					}

				}
				v = append(v, vi)
			}
		}
		return nil
	})
	return
}

//func (this *Session) Map() (v []map[string]interface{}, err error) {
//	err = this.AutoClose(func() error {
//		v, err = this.sess.SQL(this.SelectSql(nil, true), this.args...).QueryInterface()
//		return err
//	})
//	return
//}

func (this *Session) Exists() (v bool, err error) {
	err = this.AutoClose(func() error {
		v, err = this.sess.SQL(this.SelectSql(nil, false, "1")+" limit 1", this.args...).Exist()
		return err
	})
	return
}
