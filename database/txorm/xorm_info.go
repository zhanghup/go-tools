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

func (this *Engine) Tables() ([]Table, error) {
	tabs, err := this.DB.DBMetas()
	if err != nil {
		return nil, err
	}

	result := make([]Table, 0)

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

	return result, nil
}

func (this *Engine) Table(name string) (Table, error) {
	if this.tables == nil || len(this.tables) == 0 {
		_, err := this.Tables()
		if err != nil {
			return Table{}, err
		}
	}

	for _, o := range this.tables {
		if o.Name == name {
			return o, nil
		}
	}
	return Table{}, nil
}

func (this Table) Column(name string) Column {
	for _, o := range this.Columns {
		if o.Name == name {
			return o
		}
	}
	return Column{}
}
