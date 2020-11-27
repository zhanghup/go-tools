package txorm

func (this *Session) Update(bean interface{}, condiBean ...interface{}) error {
	_, err := this.Sess.Update(bean, condiBean...)
	return err
}
