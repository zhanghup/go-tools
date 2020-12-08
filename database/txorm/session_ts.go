package txorm

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
		err := this.Commit()
		if err != nil {
			return err
		}
		return this.Close()
	}
	return nil
}
