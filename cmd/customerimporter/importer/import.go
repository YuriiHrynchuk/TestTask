package importer

import (
	"encoding/csv"
	"fmt"
	"os"
)

type CustomerImporter interface {
	GetAllCustomers(filePath string) (*os.File, error)
}

type CSVImport struct{}

func NewCSVProcessor() *CSVImport {
	return &CSVImport{}
}

func (c *CSVImport) GetAllCustomers(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Problem is here (GetAllCustomers)")
		return nil, fmt.Errorf("failed to open file: %v", err)
	}

	reader := csv.NewReader(file)

	var records [][]string
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}
		records = append(records, record)
	}

	defer file.Close()

	return records, err
}
