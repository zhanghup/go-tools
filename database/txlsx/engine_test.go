package extraction

import (
	"fmt"
	"os"
	"testing"
)

func TestEngineCSV(t *testing.T) {
	f, err := os.Open("./engine_csv.csv")
	if err != nil {
		t.Error(err)
		return
	}

	res, err := NewExtractionIO(f, nil)
	if err != nil {
		t.Error(err)
		return
	}
	for _, s := range res {
		for _, row := range s.DataMap() {
			fmt.Println(row.Data())
		}
	}
}

func TestEngineXLS(t *testing.T) {
	f, err := os.Open("./engine_xls.xls")
	if err != nil {
		t.Error(err)
		return
	}

	res, err := NewExtractionIO(f, nil)
	if err != nil {
		t.Error(err)
		return
	}
	for _, s := range res {
		for _, row := range s.DataMap() {
			fmt.Println(row.Data())
		}
	}
}

func TestEngineXLSX(t *testing.T) {
	f, err := os.Open("./engine_xlsx.xlsx")
	if err != nil {
		t.Error(err)
		return
	}

	res, err := NewExtractionIO(f, nil)
	if err != nil {
		t.Error(err)
		return
	}
	for _, s := range res {
		for _, row := range s.DataMap() {
			fmt.Println(row.Data())
		}
	}
}
