package customerimporter

import (
	"TeamworkTestTask/cmd/customerimporter/domaincounter"
	"TeamworkTestTask/cmd/customerimporter/importer"
	"TeamworkTestTask/cmd/customerimporter/sorter"
	"TeamworkTestTask/internal/variables"
	"bytes"
	"fmt"
	sort2 "sort"
	"strings"
)

func CustomerProcessor() {
	processor := importer.NewCSVProcessor()
	sort := sorter.NewCSVSort()
	count := domaincounter.NewCSVCounter()

	columnInx := 0

	file, err := processor.GetAllCustomers(variables.FilePath)
	if err != nil {
		fmt.Println("сustomer_importer -> customerProcessor -> GetAllCustomers, err: ", err)
	}

	sortedFile, err := sort.SortCSV(file, columnInx)
	if err != nil {
		fmt.Println("сustomer_importer -> customerProcessor -> SortCSV, err: ", err)
	}

	var buffer bytes.Buffer

	buffer.WriteString("Sorted CSV:\n")
	for _, record := range sortedFile {
		buffer.WriteString(strings.Join(record, ", ") + "\n")
	}

	domainCount, err := count.CountDomains(file)

	var domains []string
	for domain := range domainCount {
		domains = append(domains, domain)
	}
	sort2.Strings(domains)

	buffer.WriteString("\nDomain Counts:\n")
	for _, domain := range domains {
		buffer.WriteString(fmt.Sprintf("%s: %d\n", domain, domainCount[domain]))
	}

	fmt.Print(buffer.String())
}
