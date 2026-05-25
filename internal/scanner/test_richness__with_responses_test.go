//ff:func feature=scan type=extract control=sequence
//ff:what TestRichness_WithResponses 테스트
package scanner

import "testing"

func TestRichness_WithResponses(t *testing.T) {
	ep := Endpoint{
		Responses: []Response{{TypeName: "Resp", Fields: []Field{{Name: "A"}}}},
	}
	if richness(ep) != 3 { // 2 for TypeName + 1 for field
		t.Fatalf("expected 3, got %d", richness(ep))
	}
}
