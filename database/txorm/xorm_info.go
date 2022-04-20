package txorm

type Tables []Table

type Table struct {
	Name string `json:"name"`

	Columns []Column `json:"columns"`
}

type Column struct {
	Name           string `json:"name"`
	SQLType        string `json:"sql_type"`
	SQLTypeLength1 int    `json:"sql_type_length1"`
	SQLTypeLength2 int    `json:"sql_type_length2"`
}

func (this *Engine) Tables() []Table {
	result := make([]Table, 0)

	tabs, err := this.DB.DBMetas()
	if err != nil {
		return result
	}

	for _, o := range tabs {

		columns := make([]Column, 0)
		for _, oo := range o.Columns() {
			columns = append(columns, Column{
				Name:           oo.Name,
				SQLType:        oo.SQLType.Name,
				SQLTypeLength1: oo.SQLType.DefaultLength,
				SQLTypeLength2: oo.SQLType.DefaultLength2,
			})
		}

		result = append(result, Table{
			Name:    o.Name,
			Columns: columns,
		})
	}

	this.tables = result
	return result
}

func (this *Engine) Table(name string) Table {
	if this.tables == nil || len(this.tables) == 0 {
		this.Tables()
	}

	for _, o := range this.tables {
		if o.Name == name {
			return o
		}
	}
	return Table{}
}

func (this Table) Column(name string) Column {
	for _, o := range this.Columns {
		if o.Name == name {
			return o
		}
	}
	return Column{}
}

func (this Table) ColumnExist(name string) bool {
	for _, o := range this.Columns {
		if o.Name == name {
			return true
		}
	}
	return false
}

func (this *Engine) TableColumnExist(table, column string) bool {
	for _, tab := range this.Tables() {
		if tab.Name == table {
			return tab.ColumnExist(column)
		}
	}
	return false
}

func (this *Engine) DropTables(beans ...any) error {
	return this.DB.DropTables(beans...)
}
