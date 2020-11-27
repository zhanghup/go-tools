package txorm

func (this *Session) Delete(bean interface{}) error {
	_, err := this.Sess.Delete(bean)
	return err
}
