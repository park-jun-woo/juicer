//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestExtractRolesAllowed 테스트
package quarkus

import "testing"

func TestExtractRolesAllowed(t *testing.T) {
	root, _ := parseJava([]byte(`@RolesAllowed({"admin"}) class R {}`))
	src := []byte(`@RolesAllowed({"admin"}) class R {}`)
	cls := findAllByType(root, "class_declaration")[0]
	roles := extractRolesAllowed(cls, src)
	if len(roles) != 1 || roles[0] != "admin" {
		t.Fatalf("got %v", roles)
	}
}
