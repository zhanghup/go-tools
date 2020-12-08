package txorm

func (this *Session) Insert(bean ...interface{}) error {
	_, err := this.sess.Insert(bean...)
	return err
}
