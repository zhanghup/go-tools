package extraction

type ExcelDictItem struct {
	Code  string `json:"code" xorm:"code"`
	Name  string `json:"name" xorm:"name"`
	Value string `json:"value" xorm:"value"`
}

type ExcelCustomList []ExcelCustomItem
type ExcelCustomItem struct {
	Value string `json:"value" xorm:"value"`
	Name  string `json:"name" xorm:"name"`
}

type Config struct {
	dictmap map[string][]ExcelDictItem
	custom  map[string]ExcelCustomList

	typeMap map[string]CellType
}

func NewConfig() *Config {
	return &Config{
		dictmap: map[string][]ExcelDictItem{},
		custom:  map[string]ExcelCustomList{},
		typeMap: map[string]CellType{},
	}
}

// CellTypeMap 设置Sheet中Cell的值类型
func (this *Config) CellTypeMap(m map[string]CellType) *Config {
	this.typeMap = m
	return this
}

// GetDictPtrName 根据字典值获取字典名称
func (this *Config) GetDictPtrName(key, val string) *string {
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

// GetDictName 根据字典值获取字典名称
func (this *Config) GetDictName(key, val string) string {
	v := this.GetDictPtrName(key, val)
	if v == nil {
		return ""
	}
	return *v
}

// GetDictPtrValue 根据字典名称获取字典值
func (this *Config) GetDictPtrValue(key, name string) *string {
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

// GetDictValue 根据字典名称获取字典值
func (this *Config) GetDictValue(key, name string) string {
	v := this.GetDictPtrValue(key, name)
	if v == nil {
		return ""
	}
	return *v
}

// GetCustomPtrValue 根据名称获取值
func (this *Config) GetCustomPtrValue(key, name string) *string {
	ls, ok := this.custom[key]
	if !ok {
		return nil
	}
	for _, o := range ls {
		if o.Name == name {
			return &o.Value
		}
	}
	return nil
}

// GetCustomValue 根据名称获取值
func (this *Config) GetCustomValue(key, name string) string {
	v := this.GetCustomPtrValue(key, name)
	if v == nil {
		return ""
	}
	return *v
}

// GetCustomPtrName 根据值获取名称
func (this *Config) GetCustomPtrName(key, value string) *string {
	ls, ok := this.custom[key]
	if !ok {
		return nil
	}
	for _, o := range ls {
		if o.Value == value {
			return &o.Name
		}
	}
	return nil
}

// GetCustomName 根据值获取名称
func (this *Config) GetCustomName(key, value string) string {
	v := this.GetCustomPtrName(key, value)
	if v == nil {
		return ""
	}
	return *v
}

// SetDicts 设置字典值
func (this *Config) SetDicts(dicts []ExcelDictItem) {
	this.dictmap = map[string][]ExcelDictItem{}
	for _, o := range dicts {
		if _, ok := this.dictmap[o.Code]; ok {
			this.dictmap[o.Code] = append(this.dictmap[o.Code], o)
		} else {
			this.dictmap[o.Code] = []ExcelDictItem{o}
		}
	}
}

// SetCustom 设置自定义字典值
func (this *Config) SetCustom(name string, item []ExcelCustomItem) {
	this.custom[name] = item
}
