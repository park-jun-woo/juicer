package scanner

import "testing"

func TestStatusDescription_200(t *testing.T) {
	if statusDescription("200") != "OK" {
		t.Fatal("expected OK")
	}
}

func TestStatusDescription_404(t *testing.T) {
	if statusDescription("404") != "Not Found" {
		t.Fatal("expected Not Found")
	}
}

func TestStatusDescription_Unknown(t *testing.T) {
	if statusDescription("(unknown)") != "Error" {
		t.Fatal("expected Error")
	}
}

func TestStatusDescription_Other(t *testing.T) {
	if statusDescription("418") != "Response" {
		t.Fatal("expected Response")
	}
}
