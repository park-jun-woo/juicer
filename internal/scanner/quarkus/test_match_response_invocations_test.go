//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestMatchResponseInvocations 테스트
package quarkus

import "testing"

func TestMatchResponseInvocations(t *testing.T) {
	root, _ := parseJava([]byte(`class R { void m() { return Response.status(204).build(); } }`))
	src := []byte(`class R { void m() { return Response.status(204).build(); } }`)
	body := findAllByType(root, "block")[0]
	if code := matchResponseInvocations(body, src); code != "204" {
		t.Fatalf("got %q", code)
	}
}
