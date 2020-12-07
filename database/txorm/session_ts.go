package txorm

func (this *Session) TS(fn func(sess ISession) error) error {
	err := this.Begin()
	if err != nil {
		return err
	}
	err = fn(this)
	if err != nil {
		_ =this.Rollback()
		return err
	}
	return this.Commit()
}
