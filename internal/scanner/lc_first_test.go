package scanner

import "testing"

func TestLcFirst_Simple(t *testing.T) {
	if lcFirst("Building") != "building" {
		t.Fatal("expected building")
	}
}

func TestLcFirst_Acronym(t *testing.T) {
	if lcFirst("ID") != "id" {
		t.Fatal("expected id")
	}
}

func TestLcFirst_SMSResult(t *testing.T) {
	if lcFirst("SMSResult") != "smsResult" {
		t.Fatal("expected smsResult")
	}
}

func TestLcFirst_Empty(t *testing.T) {
	if lcFirst("") != "" {
		t.Fatal("expected empty")
	}
}

func TestLcFirst_Lowercase(t *testing.T) {
	if lcFirst("already") != "already" {
		t.Fatal("expected already")
	}
}
