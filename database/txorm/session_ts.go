package txorm

func (this *Session) TS(fn func(sess *Session) error) error {
	err := this.Sess.Begin()
	if err != nil {
		panic(err)
	}
	err = fn(this)
	if err != nil {
		err2 := this.Sess.Rollback()
		if err2 != nil {
			panic(err2)
		}
		return err
	}

	if this.context == nil {
		err2 := this.Sess.Commit()
		if err2 != nil {
			panic(err2)
		}
		if this.autoClose {
			this.Sess.Close()
		}
	} else {
		go func() {
			<-this.context.Done()
			err2 := this.Sess.Commit()
			if err2 != nil {
				panic(err2)
			}
			this.Sess.Close()
		}()
	}

	return nil

}
