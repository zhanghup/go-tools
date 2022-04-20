package buntdb

func (this *Engine) Ts(fn func(sess ISession) error) error {
	tx, err := NewSession(this.db)
	if err != nil {
		return err
	}
	ttx := tx.(*Session)
	err = fn(tx)
	if err != nil {
		_ = ttx.Rollback()
		return err
	}
	return ttx.Commit()
}
