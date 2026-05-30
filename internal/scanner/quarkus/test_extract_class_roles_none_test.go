//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestExtractClassRoles_None 테스트
package quarkus

import "testing"

func TestExtractClassRoles_None(t *testing.T) {
	root, _ := parseJava([]byte(`class R {}`))
	src := []byte(`class R {}`)
	cls := findAllByType(root, "class_declaration")[0]
	if roles := extractClassRoles(cls, src); roles != nil {
		t.Fatalf("got %v", roles)
	}
}
