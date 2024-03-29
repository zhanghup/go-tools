package txorm

import (
	"fmt"
	"github.com/zhanghup/go-tools"
	"strings"
	"time"
	"xorm.io/xorm/core"
)

func (this *Session) SelectSql(bean any, orderFlag bool, columns ...string) string {
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

func (this *Session) Find(bean any) error {
	return this.AutoClose(func() error {
		return this.sess.SQL(this.SelectSql(bean, true), this.args...).Find(bean)
	})
}

func (this *Session) Get(bean any) (v bool, err error) {
	err = this.AutoClose(func() error {
		v, err = this.sess.SQL(this.SelectSql(bean, true)+" limit 1", this.args...).Get(bean)
		return err
	})
	return
}

func (this *Session) Map() (v []map[string]any, err error) {
	err = this.AutoClose(func() error {
		var rows *core.Rows
		if this.autoClose {
			rows, err = this.sess.DB().Query(this.SelectSql(nil, true), this.args...)
		} else {
			rows, err = this.sess.Tx().Query(this.SelectSql(nil, true), this.args...)
		}

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
				vi := map[string]any{}

				for _, o := range types {
					if colValue, ok := vv[o.Name()]; ok && colValue != nil {
						newValue := string(colValue)

						switch o.DatabaseTypeName() {
						case "DATETIME":
							if this._db.DriverName() == "sqlite3" {
								vi[o.Name()], err = time.ParseInLocation("2006-01-02T15:04:05Z", newValue, time.Local)
								if err != nil {
									return err
								}
							} else {
								vi[o.Name()], err = time.ParseInLocation("2006-01-02 15:04:05", newValue, time.Local)
								if err != nil {
									return err
								}
							}

						case "BOOL", "TINYINT":
							vi[o.Name()] = tools.StrToInt[int8](newValue)
						case "BLOB":
							vi[o.Name()] = colValue
						case "FLOAT":
							vi[o.Name()] = tools.StrToFloat[float32](newValue)
						case "DOUBLE", "REAL":
							vi[o.Name()] = tools.StrToFloat[float64](newValue)
						case "BIGINT":
							vi[o.Name()] = tools.StrToInt[int64](newValue)
						case "INT", "INTEGER":
							vi[o.Name()] = tools.StrToInt[int](newValue)
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

func (this *Session) MapString() (v []map[string]string, err error) {
	err = this.AutoClose(func() error {
		var rows *core.Rows
		if this.autoClose {
			rows, err = this.sess.DB().Query(this.SelectSql(nil, true), this.args...)
		} else {
			rows, err = this.sess.Tx().Query(this.SelectSql(nil, true), this.args...)
		}

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
				vi := map[string]string{}

				for _, o := range types {
					if colValue, ok := vv[o.Name()]; ok && colValue != nil {
						vi[o.Name()] = string(colValue)
					} else {
						vi[o.Name()] = ""
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
