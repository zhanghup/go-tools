package commonxl

import (
	"strings"
	"time"
)

// ConvertToDate converts a floating-point value using the
// Excel date serialization conventions.
func (x *Formatter) ConvertToDate(val float64) time.Time {
	// http://web.archive.org/web/20190808062235/http://aa.usno.navy.mil/faq/docs/JD_Formula.php
	v := int(val)
	if v < 61 {
		jdate := val + 0.5
		if (x.flags & fMode1904) != 0 {
			jdate += 2416480.5
		} else {
			jdate += 2415018.5
		}
		JD := int(jdate)
		frac := jdate - float64(JD)

		L := JD + 68569
		N := 4 * L / 146097
		L = L - (146097*N+3)/4
		I := 4000 * (L + 1) / 1461001
		L = L - 1461*I/4 + 31
		J := 80 * L / 2447
		day := L - 2447*J/80
		L = J / 11
		month := time.Month(J + 2 - 12*L)
		year := 100*(N-49) + I + L

		t := time.Duration(float64(time.Hour*24) * frac)
		return time.Date(year, month, day, 0, 0, 0, 0, time.UTC).Add(t)
	}
	frac := val - float64(v)
	date := time.Date(1904, 1, 1, 0, 0, 0, 0, time.UTC)
	if (x.flags & fMode1904) == 0 {
		date = time.Date(1899, 12, 30, 0, 0, 0, 0, time.UTC)
	}

	t := time.Duration(float64(time.Hour*24) * frac)
	return date.AddDate(0, 0, v).Add(t)
}

func timeFmtFunc(f string) FmtFunc {
	return func(x *Formatter, v interface{}) string {
		t, ok := v.(time.Time)
		if !ok {
			fval, ok := convertToFloat64(v)
			if !ok {
				return "MUST BE time.Time OR numeric TO FORMAT CORRECTLY"
			}
			t = x.ConvertToDate(fval)
		}
		//log.Println("formatting date", t, "with", f, "=", t.Format(f))
		return t.Format(f)
	}
}

// same as above but replaces "AM" and "PM" with chinese translations.
// TODO: implement others
func cnTimeFmtFunc(f string) FmtFunc {
	return func(x *Formatter, v interface{}) string {
		t, ok := v.(time.Time)
		if !ok {
			fval, ok := convertToFloat64(v)
			if !ok {
				return "MUST BE time.Time OR numeric TO FORMAT CORRECTLY"
			}
			t = x.ConvertToDate(fval)
		}
		s := t.Format(f)
		s = strings.Replace(s, `AM`, `上午`, 1)
		return strings.Replace(s, `PM`, `下午`, 1)
	}
}
