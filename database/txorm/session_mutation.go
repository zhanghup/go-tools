package txorm

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
	this.Begin()
	err := fn(this)
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

func (this *Session) Exec() error {
	if this.autoClose {
		// 由engine直接进入的方法，需要自动关闭session
		defer this.AutoClose()
	}

	sqls := []interface{}{this._sql_with() + " " + this._sql()}
	_, err := this.sess.Exec(append(sqls, this.args...)...)
	return err
}
