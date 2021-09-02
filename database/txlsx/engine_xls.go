package extraction

import (
	"github.com/zhanghup/go-tools/database/txlsx/xls"
	"io"
	"os"
	"strings"
)

type EngineXls struct {
	Config *Config
}

func NewEngineXls(cfg *Config) IExtraction {
	return &EngineXls{
		Config: cfg,
	}
}

// Open 读取csv文件数据
// filename 文件路径
// cfg 配置文件，可空可接收一个配置参数
func (this *EngineXls) Open(filename string) ([]Sheet, error) {
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
func (this *EngineXls) OpenIO(read io.Reader) ([]Sheet, error) {
	shs, err := xls.OpenIO(read)
	if err != nil {
		return nil, err
	}
	sheets, err := shs.List()
	if err != nil {
		return nil, err
	}

	result := make([]Sheet, 0)

	for _, name := range sheets {
		sheet, err := shs.Get(name)
		if err != nil {
			return nil, err
		}

		rsheet := NewSheet(this.Config)
		rsheet.Name = name

		for sheet.Next() {
			items := make([]Cell, 0)
			cells := sheet.Strings()
			types := sheet.Types()
			fmts := sheet.Formats()

			for i := 0; i < len(cells); i++ {
				ty := this.Type(types[i], fmts[i])
				items = append(items, Cell{
					Config:    this.Config,
					Value:     cells[i],
					Type:      ty,
					Formatter: this.Formatter(ty, fmts[i]),
				})
			}
			rsheet.DataList = append(rsheet.DataList, items)
		}

		result = append(result, *rsheet)
	}

	return result, nil
}

func (this *EngineXls) Formatter(t CellType, fmt string) string {
	switch t {
	case CellTypeDate:
		return fmt
	default:
		return ""
	}
}

// Type "boolean", "integer", "float", "string", "date",
func (this *EngineXls) Type(t string, fmt string) CellType {
	// 判断格式化类型是否未日期类型
	if fmt != "General" {
		s := []string{"ss", "mm", "h", "d", "m", "yyyy", "yy"}
		n := 0
		for _, key := range s {
			if strings.Index(fmt, key) != -1 {
				n += 1
			}
		}
		if n > 2 {
			return CellTypeDate
		}
	}

	switch t {
	case "integer":
		return CellTypeInt
	case "float":
		return CellTypeFloat
	case "date":
		return CellTypeDate
	default:
		return CellTypeString
	}
}
