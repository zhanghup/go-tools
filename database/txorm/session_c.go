package txorm

func (this *Session) Insert(bean ...interface{}) error {
	_, err := this.Sess.Insert(bean...)
	return err
}
