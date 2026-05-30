//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestMatchBodyInvocations 테스트
package spring

import "testing"

func TestMatchBodyInvocations(t *testing.T) {
	root, src := parseS(t, `class C { void m() { return ResponseEntity.noContent().build(); } }`)
	body := findAllByType(root, "block")[0]
	if code := matchBodyInvocations(body, src); code != "204" {
		t.Fatalf("got %q", code)
	}
}
