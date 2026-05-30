//ff:func feature=scan type=test control=iteration dimension=1 topic=quarkus
//ff:what TestMatchResponseInvocation_Status 테스트
package quarkus

import "testing"

func TestMatchResponseInvocation_Status(t *testing.T) {
	root, _ := parseJava([]byte(`class R { void m() { Response.status(404).build(); } }`))
	src := []byte(`class R { void m() { Response.status(404).build(); } }`)
	invs := findAllByType(root, "method_invocation")
	for _, inv := range invs {
		if code := matchResponseInvocation(inv, src); code == "404" {
			return
		}
	}
	t.Fatal("did not match 404")
}
