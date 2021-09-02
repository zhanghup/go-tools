package extraction

import (
	"encoding/csv"
	"io"
	"os"
	"regexp"
)

type EngineCsv struct {
	Config *Config
}

func NewEngineCsv(cfg *Config) IExtraction {
	return &EngineCsv{Config: cfg}
}

// Open 读取csv文件数据
// filename 文件路径
// cfg 配置文件，可空可接收一个配置参数
func (this *EngineCsv) Open(filename string) ([]Sheet, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return this.OpenIO(f)
}

// OpenIO 读取csv文件数据
// read 文件流
// cfg 配置文件，可空可接收一个配置参数
func (this *EngineCsv) OpenIO(read io.Reader) ([]Sheet, error) {
	excel := csv.NewReader(read)
	data, err := excel.ReadAll()
	if err != nil {
		return nil, err
	}

	sh := NewSheet(this.Config)
	sh.Name = "sheet1"

	for _, row := range data {
		items := make([]Cell, 0)
		for _, cel := range row {
			items = append(items, Cell{
				Config:    sh.Config,
				Type:      CellTypeString,
				Value:     cel,
				Formatter: "",
			})
		}
		sh.DataList = append(sh.DataList, items)
	}

	return []Sheet{*sh}, nil
}

var engineCsvFloat = regexp.MustCompile(`^\d+\.\d+$`)
var engineCsvInt = regexp.MustCompile(`^\d+$`)

func (this *EngineCsv) CellType(val string) CellType {
	switch {
	case engineCsvInt.MatchString(val):
		return CellTypeInt
	case engineCsvFloat.MatchString(val):
		return CellTypeFloat
	default:
		return CellTypeString
	}
}
