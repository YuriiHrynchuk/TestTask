package test

import (
	"TeamworkTestTask/cmd/customerimporter/sorter"
	"reflect"
	"testing"
)

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
