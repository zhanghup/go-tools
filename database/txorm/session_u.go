package txorm

func (this *Session) Update(bean interface{}, condiBean ...interface{}) error {
	_, err := this.sess.Update(bean, condiBean...)
	return err
}
