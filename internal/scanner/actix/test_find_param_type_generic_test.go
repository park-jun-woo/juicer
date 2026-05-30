//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestFindParamType_Generic 테스트
package actix

import "testing"

func TestFindParamType_Generic(t *testing.T) {
	src := []byte(`fn f(body: web::Json<User>) {}`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	p := firstParam(root)
	if p == nil {
		t.Fatal("no parameter")
	}
	ty := findParamType(p)
	if ty == nil {
		t.Fatal("expected a type node")
	}
	if ty.Type() != "generic_type" {
		t.Errorf("type kind = %s, want generic_type", ty.Type())
	}
}
