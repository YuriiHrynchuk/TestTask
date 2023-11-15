package domaincounter

import (
	"fmt"
	"os"
)

type DomainCounter interface {
	CountDomains(file *os.File) (map[string]int, error)
}

type CSVCounter struct{}

func NewCSVCounter() *CSVCounter {
	return &CSVCounter{}
}

func (c *CSVCounter) CountDomains(records [][]string) (map[string]int, error) {
	if records == nil {
		return nil, fmt.Errorf("File is empty")
	}

	domainCount := make(map[string]int)

	for _, record := range records {
		email := record[2]
		domain := ExtractDomain(email)
		domainCount[domain]++
	}

	return domainCount, nil
}
