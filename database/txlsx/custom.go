package txlsx

type ExcelDictItem struct {
	Code  string `json:"code" xorm:"code"`
	Name  string `json:"name" xorm:"name"`
	Value string `json:"value" xorm:"value"`
}

type ExcelCustomList []ExcelCustomItem
type ExcelCustomItem struct {
	Id   string `json:"id" xorm:"id"`
	Name string `json:"name" xorm:"name"`
}

type ExcelExt struct {
	dictmap map[string][]ExcelDictItem
	custom  map[string]ExcelCustomList
}

func (this ExcelExt) GetDictPtrName(key, val string) *string {
	dicts, ok := this.dictmap[key]
	if !ok {
		return nil
	}

	for _, o := range dicts {
		if o.Value == val {
			return &o.Name
		}
	}
	return nil
}

func (this ExcelExt) GetDictName(key, val string) string {
	v := this.GetDictPtrName(key, val)
	if v == nil {
		return ""
	}
	return *v
}

func (this ExcelExt) GetDictPtrValue(key, name string) *string {
	dicts, ok := this.dictmap[key]
	if !ok {
		return nil
	}

	for _, o := range dicts {
		if o.Code == key && o.Name == name {
			return &o.Value
		}
	}
	return nil
}

func (this ExcelExt) GetDictValue(key, name string) string {
	v := this.GetDictPtrValue(key, name)
	if v == nil {
		return ""
	}
	return *v
}

func (this ExcelExt) GetCustomPtrId(key, name string) *string {
	ls, ok := this.custom[key]
	if !ok {
		return nil
	}
	for _, o := range ls {
		if o.Name == name {
			return &o.Id
		}
	}
	return nil
}

func (this ExcelExt) GetCustomId(key, name string) string {
	v := this.GetCustomPtrId(key, name)
	if v == nil {
		return ""
	}
	return *v
}

func (this ExcelExt) GetCustomPtrName(key, id string) *string {
	ls, ok := this.custom[key]
	if !ok {
		return nil
	}
	for _, o := range ls {
		if o.Id == id {
			return &o.Name
		}
	}
	return nil
}

func (this ExcelExt) GetCustomName(key, id string) string {
	v := this.GetCustomPtrName(key, id)
	if v == nil {
		return ""
	}
	return *v
}
