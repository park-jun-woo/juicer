//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestExtractClassRoles_Authenticated 테스트
package quarkus

import "testing"

func TestExtractClassRoles_Authenticated(t *testing.T) {
	root, _ := parseJava([]byte(`@Authenticated class R {}`))
	src := []byte(`@Authenticated class R {}`)
	cls := findAllByType(root, "class_declaration")[0]
	roles := extractClassRoles(cls, src)
	if len(roles) != 1 || roles[0] != "**" {
		t.Fatalf("got %v", roles)
	}
}
