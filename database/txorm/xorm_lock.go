package txorm

func (this *Engine) lock() {
	if this.DB.DriverName() == "sqlite3" {
		this.sqliteSync.Lock()
	}
}

func (this *Engine) unlock() {
	if this.DB.DriverName() == "sqlite3" {
		this.sqliteSync.Unlock()
	}
}
