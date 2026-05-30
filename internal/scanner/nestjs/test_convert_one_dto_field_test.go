//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestConvertOneDtoField 테스트
package nestjs

import "testing"

func TestConvertOneDtoField(t *testing.T) {
	df := dtoField{name: "age", tsType: "number", optional: false, validators: []string{"min:0"}}
	sf := convertOneDtoField(df)
	if sf.Name != "age" {
		t.Fatalf("got %+v", sf)
	}
	if sf.Validate == "" {
		t.Fatalf("expected validate built, got %+v", sf)
	}
}
