//ff:func feature=scan type=test control=sequence
//ff:what TestRichness_RichCov 테스트
package scanner

import "testing"

func TestRichness_RichCov(t *testing.T) {
	ep := Endpoint{
		Request: &Request{Body: &Body{TypeName: "Req", Fields: []Field{{Name: "x"}}}},
		Responses: []Response{
			{TypeName: "Resp", Fields: []Field{{Name: "id"}}},
		},
	}
	if richness(ep) < 5 {
		t.Fatalf("expected >= 5, got %d", richness(ep))
	}
}
