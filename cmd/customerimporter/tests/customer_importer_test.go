package tests

import (
	"TeamworkTestTask/cmd/customerimporter"
	"TeamworkTestTask/cmd/customerimporter/domaincounter"
	"TeamworkTestTask/cmd/customerimporter/importer"
	"TeamworkTestTask/cmd/customerimporter/sorter"
	"TeamworkTestTask/internal/variables"
	"TeamworkTestTask/pkg/domainextract"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestGetAllCustomers(t *testing.T) {
	testData := "Justin,Hansen,jhansen3@360.cn,Male,251.166.224.119"
	tmpFile := createTempCSVFile(testData)
	defer removeTempCSVFile(tmpFile)

	csvImport := importer.NewCSVProcessor()
	records, err := csvImport.GetAllCustomers(tmpFile)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expectedRecords := [][]string{
		{"Justin", "Hansen", "jhansen3@360.cn", "Male", "251.166.224.119"},
	}

	if !reflect.DeepEqual(records, expectedRecords) {
		t.Errorf("Records do not match.\nExpected: %v\nGot: %v", expectedRecords, records)
	}
}

func TestCustomerProcessorWithInvalidFilePath(t *testing.T) {
	invalidFilePath := "invalid/file/path.csv"
	variables.FilePath = invalidFilePath

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	customerimporter.CustomerProcessor()

	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	w.Close()
	os.Stdout = oldStdout
	out := <-outC

	expectedErrorMessage := fmt.Sprintf("failed to open file: open %s: no such file or directory", invalidFilePath)

	if !strings.Contains(out, expectedErrorMessage) {
		t.Errorf("Expected error message for invalid file path, but got:\n%s", out)
	}
}

func createTempCSVFile(data string) string {
	tmpFile, err := os.CreateTemp("", "testfile*.csv")
	if err != nil {
		log.Fatal(err)
	}

	if _, err := tmpFile.Write([]byte(data)); err != nil {
		log.Fatal(err)
	}

	return tmpFile.Name()
}

func TestSortCSV(t *testing.T) {
	testRecords := [][]string{
		{"Justin", "Hansen", "jhansen3@360.cn", "Male", "251.166.224.119"},
		{"Lillian", "Lawrence", "llawrence9@blogtalkradio.com", "Female", "190.106.124.105"},
		{"Ernest", "Reid", "ereid5@rediff.com", "Male", "243.219.170.46"},
		{"Bonnie", "Ortiz", "bortiz1@cyberchimps.com", "Female", "197.54.209.129"},
	}

	csvSort := sorter.NewCSVSort()
	sortedRecords, err := csvSort.SortCSV(testRecords, 0)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expectedSortedRecords := [][]string{
		{"Bonnie", "Ortiz", "bortiz1@cyberchimps.com", "Female", "197.54.209.129"},
		{"Ernest", "Reid", "ereid5@rediff.com", "Male", "243.219.170.46"},
		{"Justin", "Hansen", "jhansen3@360.cn", "Male", "251.166.224.119"},
		{"Lillian", "Lawrence", "llawrence9@blogtalkradio.com", "Female", "190.106.124.105"},
	}

	if !reflect.DeepEqual(sortedRecords, expectedSortedRecords) {
		t.Errorf("Sorted records do not match.\nExpected: %v\nGot: %v", expectedSortedRecords, sortedRecords)
	}
}

func TestCountDomains(t *testing.T) {
	testRecords := [][]string{
		{"Justin", "Hansen", "jhansen3@360.cn", "Male", "251.166.224.119"},
		{"Lillian", "Lawrence", "llawrence9@blogtalkradio.com", "Female", "190.106.124.105"},
		{"Ernest", "Reid", "ereid5@rediff.com", "Male", "243.219.170.46"},
		{"Bonnie", "Ortiz", "bortiz1@cyberchimps.com", "Female", "197.54.209.129"},
	}

	csvCounter := domaincounter.NewCSVCounter()
	domainCount, err := csvCounter.CountDomains(testRecords)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expectedDomainCount := map[string]int{
		"360.cn":            1,
		"blogtalkradio.com": 1,
		"rediff.com":        1,
		"cyberchimps.com":   1,
	}

	if !reflect.DeepEqual(domainCount, expectedDomainCount) {
		t.Errorf("Domain count does not match.\nExpected: %v\nGot: %v", expectedDomainCount, domainCount)
	}
}

func TestExtractDomain(t *testing.T) {
	testEmail := "bortiz1@cyberchimps.com"

	domain := domainextract.ExtractDomain(testEmail)

	expectedDomain := "cyberchimps.com"

	if domain != expectedDomain {
		t.Errorf("Extracted domain does not match.\nExpected: %v\nGot: %v", expectedDomain, domain)
	}
}

func removeTempCSVFile(filePath string) {
	if err := os.Remove(filePath); err != nil {
		log.Fatal(err)
	}
}
