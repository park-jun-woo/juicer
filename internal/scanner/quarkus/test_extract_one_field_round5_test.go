//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestExtractOneField_Round5 테스트
package quarkus

import "testing"

func TestExtractOneField_Round5(t *testing.T) {
	root, src := qParse(t, `class Dto { public String name; }`)
	field := qFirst(t, root, "field_declaration")
	f := extractOneField(field, src)
	if f.Name != "name" {
		t.Fatalf("field: %+v", f)
	}
}
