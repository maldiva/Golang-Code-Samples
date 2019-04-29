package main

import (
	"testing"
)

func TestFoundAtSite(t *testing.T) {

	var r Request
	r.Site = []string{"https://google.com", "https://yahoo.com"}
	r.SearchText = "google"

	result, status := FoundAtSite(&r)
	if status != true && result == "" {
		t.Error("Haven't found 'google' at Google.com")
	}
}

func TestFoundAtSiteOnCorruptedRequest(t *testing.T) {

	var r Request
	r.Site = []string{"SomeTrash.123"}
	r.SearchText = "google"

	result, status := FoundAtSite(&r)
	if status != false && result != "" {
		t.Error("Incorrect URL passed the test")
	}
}

func TestTableFoundAtSite(t *testing.T) {
	var tests = []struct {
		input          Request
		expectedSite   string
		expectedStatus bool
	}{
		{input: Request{Site: []string{"https://google.com", "https://yahoo.com"}, SearchText: "google"}, expectedSite: "https://google.com", expectedStatus: true},
		{input: Request{Site: []string{"fake request"}, SearchText: "google"}, expectedSite: "", expectedStatus: false},
		{input: Request{Site: []string{"https://ts.kg", "https://geeks.team"}, SearchText: "Новые"}, expectedSite: "https://ts.kg", expectedStatus: true},
	}

	for _, test := range tests {
		if site, status := FoundAtSite(&test.input); status != test.expectedStatus || site != test.expectedSite {
			t.Errorf("Test failed: %s request, %t expected status, %t received status,\n %s expected site, %s received site", test.input, test.expectedStatus, status, test.expectedSite, site)
		}
	}
}
