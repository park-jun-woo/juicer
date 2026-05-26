//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestTsTypeToField_Optional 테스트
package nestjs

import "testing"

func TestTsTypeToField_Optional(t *testing.T) {
	f := tsTypeToField("age", "number", true)
	if f.Validate != "optional" {
		t.Fatalf("expected optional, got %q", f.Validate)
	}
}
