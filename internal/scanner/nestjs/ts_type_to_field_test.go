//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestTsTypeToField_String 테스트
package nestjs

import "testing"

func TestTsTypeToField_String(t *testing.T) {
	f := tsTypeToField("name", "string", false)
	if f.Name != "name" || f.Type != "string" {
		t.Fatalf("unexpected: %+v", f)
	}
	if f.Validate != "" {
		t.Fatalf("expected empty validate, got %q", f.Validate)
	}
}
