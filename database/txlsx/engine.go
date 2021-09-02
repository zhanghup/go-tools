package extraction

import (
	"bytes"
	"errors"
	"io"
)

type FileType int

const (
	CSV FileType = iota
	XLS
	XLSX
)

type IExtraction interface {
	Open(filename string) ([]Sheet, error)
	OpenIO(read io.Reader) ([]Sheet, error)
}

func NewExtraction(filename string, cfg *Config, ty ...FileType) ([]Sheet, error) {
	if len(ty) > 0 {
		switch ty[0] {
		case CSV:
			return NewEngineCsv(cfg).Open(filename)
		case XLS:
			return NewEngineXls(cfg).Open(filename)
		case XLSX:
			return NewEngineXlsx(cfg).Open(filename)
		}
		return nil, errors.New("引擎不存在")
	}

	result, err := NewEngineXlsx(cfg).Open(filename)
	if err == nil {
		return result, nil
	}
	result, err = NewEngineXls(cfg).Open(filename)
	if err == nil {
		return result, nil
	}
	result, err = NewEngineCsv(cfg).Open(filename)
	if err == nil {
		return result, nil
	}

	return nil, errors.New("数据无法解析")
}

func NewExtractionIO(read io.Reader, cfg *Config, ty ...FileType) ([]Sheet, error) {
	data, err := io.ReadAll(read)
	if err != nil {
		return nil, err
	}

	if len(ty) > 0 {
		switch ty[0] {
		case CSV:
			return NewEngineCsv(cfg).OpenIO(bytes.NewBuffer(data))
		case XLS:
			return NewEngineXls(cfg).OpenIO(bytes.NewBuffer(data))
		case XLSX:
			return NewEngineXlsx(cfg).OpenIO(bytes.NewBuffer(data))
		}
		return nil, errors.New("引擎不存在")
	}

	result, err := NewEngineXlsx(cfg).OpenIO(bytes.NewBuffer(data))
	if err == nil {
		return result, nil
	}
	result, err = NewEngineXls(cfg).OpenIO(bytes.NewBuffer(data))
	if err == nil {
		return result, nil
	}
	result, err = NewEngineCsv(cfg).OpenIO(bytes.NewBuffer(data))
	if err == nil {
		return result, nil
	}

	return nil, errors.New("数据无法解析")
}
