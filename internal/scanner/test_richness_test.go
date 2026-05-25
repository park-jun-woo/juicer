//ff:func feature=scan type=convert control=sequence
//ff:what TestRichness 테스트
package scanner

import (
	"testing"
)

func TestRichness(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		ep := Endpoint{}
		if richness(ep) != 0 {
			t.Error("expected 0")
		}
	})

	t.Run("with body type name", func(t *testing.T) {
		ep := Endpoint{
			Request: &Request{
				Body: &Body{TypeName: "Req", Fields: []Field{{Name: "a"}}},
			},
		}
		if richness(ep) < 3 {
			t.Error("expected at least 3")
		}
	})

	t.Run("with response", func(t *testing.T) {
		ep := Endpoint{
			Responses: []Response{{TypeName: "Resp", Fields: []Field{{Name: "a"}}}},
		}
		if richness(ep) < 2 {
			t.Error("expected at least 2")
		}
	})
}
