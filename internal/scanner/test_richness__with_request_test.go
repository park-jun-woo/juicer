//ff:func feature=scan type=extract control=sequence
//ff:what TestRichness_WithRequest 테스트
package scanner

import "testing"

func TestRichness_WithRequest(t *testing.T) {
	ep := Endpoint{
		Request: &Request{Body: &Body{TypeName: "Req", Fields: []Field{{Name: "A"}}}},
	}
	if richness(ep) != 4 { // 3 for TypeName + 1 for field
		t.Fatalf("expected 4, got %d", richness(ep))
	}
}
