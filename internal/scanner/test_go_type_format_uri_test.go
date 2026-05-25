//ff:func feature=scan type=extract control=sequence
//ff:what TestGoTypeFormat_URI 테스트
package scanner

import "testing"

func TestGoTypeFormat_URI(t *testing.T) {
	got := goTypeFormat("string", Field{Validate: "url"})
	if got != "uri" {
		t.Fatalf("expected uri, got %s", got)
	}
}
