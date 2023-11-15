package test

import (
	"TeamworkTestTask/cmd/customerimporter/domaincounter"
	"reflect"
	"testing"
)

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

	domain := domaincounter.ExtractDomain(testEmail)

	expectedDomain := "cyberchimps.com"

	if domain != expectedDomain {
		t.Errorf("Extracted domain does not match.\nExpected: %v\nGot: %v", expectedDomain, domain)
	}
}
