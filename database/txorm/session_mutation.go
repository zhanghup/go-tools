package txorm

func (this *Session) Insert(bean ...interface{}) error {
	return this.begin(func() error {
		_, err := this.sess.Insert(bean...)
		return err
	})
}

func (this *Session) Update(bean interface{}, condiBean ...interface{}) error {
	return this.begin(func() error {
		_, err := this.sess.Update(bean, condiBean...)
		return err
	})
}

func (this *Session) Delete(bean interface{}) error {
	return this.begin(func() error {
		_, err := this.sess.Delete(bean)
		return err
	})
}

func (this *Session) TS(fn func(sess ISession) error) error {
	return this.AutoClose(func() error {
		this.Begin()
		err := fn(this)
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
		sqls := []interface{}{this._sql_with() + " " + this._sql()}
		_, err := this.sess.Exec(append(sqls, this.args...)...)
		return err
	})

}
