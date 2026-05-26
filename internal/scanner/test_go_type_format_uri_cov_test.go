//ff:func feature=scan type=test control=sequence
//ff:what TestGoTypeFormat_URICov 테스트
package scanner

import "testing"

func TestGoTypeFormat_URICov(t *testing.T) {
	f := Field{Validate: "url"}
	if got := goTypeFormat("string", f); got != "uri" {
		t.Fatalf("expected uri, got %s", got)
	}
}
