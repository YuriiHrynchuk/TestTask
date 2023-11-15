package sorter

import (
	"fmt"
	"os"
	"sort"
)

type CustomerSorter interface {
	SortCSV(file *os.File, columnIdx int) ([][]string, error)
}

type CSVSort struct{}

func NewCSVSort() *CSVSort {
	return &CSVSort{}
}

func (c *CSVSort) SortCSV(records [][]string, columnIdx int) ([][]string, error) {
	if records == nil {
		fmt.Println("Problem is here (empty file)")
		return nil, fmt.Errorf("File is empty")
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i][columnIdx] < records[j][columnIdx]
	})

	return records, nil
}
