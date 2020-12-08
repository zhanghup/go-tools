package txorm

func (this *Session) Delete(bean interface{}) error {
	_, err := this.sess.Delete(bean)
	return err
}
