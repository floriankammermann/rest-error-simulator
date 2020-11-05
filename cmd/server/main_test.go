package main

import "testing"

func TestResponseCode(t *testing.T) {
	responseCode := getResponseCode(1, 1, 200, 500)
	if responseCode != 200 {
		t.Fail()
	}
	responseCode = getResponseCode(3, 2, 200, 500)
	if responseCode != 500 {
		t.Fail()
	}
}
