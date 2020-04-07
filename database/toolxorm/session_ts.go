package toolxorm


func (this *Session) TS(fn func(sess *Session) error) {
	err := fn(this)
	if err != nil {
		err2 := this.Sess.Rollback()
		if err2 != nil {
			panic(err2)
		}
	}
	err2 := this.Sess.Commit()
	if err2 != nil {
		panic(err2)
	}
	if this.autoClose {
		this.Sess.Close()
	}
}

