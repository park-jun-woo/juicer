package scanner

import "testing"

func TestPickBestResponse_Single(t *testing.T) {
	resps := []Response{{Status: "200"}}
	best := pickBestResponse(resps)
	if best.Status != "200" {
		t.Fatal("expected 200")
	}
}

func TestPickBestResponse_Multiple(t *testing.T) {
	resps := []Response{
		{Status: "200"},
		{Status: "200", TypeName: "User", Fields: []Field{{Name: "ID"}}},
	}
	best := pickBestResponse(resps)
	if best.TypeName != "User" {
		t.Fatal("expected User (richer response)")
	}
}
