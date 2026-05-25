package scanner

import "testing"

func TestRichness_Empty(t *testing.T) {
	ep := Endpoint{}
	if richness(ep) != 0 {
		t.Fatal("expected 0")
	}
}

func TestRichness_WithRequest(t *testing.T) {
	ep := Endpoint{
		Request: &Request{Body: &Body{TypeName: "Req", Fields: []Field{{Name: "A"}}}},
	}
	if richness(ep) != 4 { // 3 for TypeName + 1 for field
		t.Fatalf("expected 4, got %d", richness(ep))
	}
}

func TestRichness_WithResponses(t *testing.T) {
	ep := Endpoint{
		Responses: []Response{{TypeName: "Resp", Fields: []Field{{Name: "A"}}}},
	}
	if richness(ep) != 3 { // 2 for TypeName + 1 for field
		t.Fatalf("expected 3, got %d", richness(ep))
	}
}
