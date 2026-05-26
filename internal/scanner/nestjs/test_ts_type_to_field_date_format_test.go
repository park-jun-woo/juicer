//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestTsTypeToField_DateFormat 테스트
package nestjs

import "testing"

func TestTsTypeToField_DateFormat(t *testing.T) {
	f := tsTypeToField("created", "Date", false)
	if f.Type != "string:date-time" {
		t.Fatalf("expected string:date-time, got %q", f.Type)
	}
}
