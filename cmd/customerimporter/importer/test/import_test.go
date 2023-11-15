package test

import (
	"TeamworkTestTask/cmd/customerimporter"
	"TeamworkTestTask/cmd/customerimporter/importer"
	"TeamworkTestTask/internal/variables"
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

func removeTempCSVFile(filePath string) {
	if err := os.Remove(filePath); err != nil {
		log.Fatal(err)
	}
}
