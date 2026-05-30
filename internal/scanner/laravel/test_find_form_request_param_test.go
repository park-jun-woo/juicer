//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestFindFormRequestParam 테스트
package laravel

import "testing"

func TestFindFormRequestParam(t *testing.T) {
	params := []methodParam{
		{name: "id", typeName: "int"},
		{name: "request", typeName: "StoreUserRequest"},
	}
	if got := findFormRequestParam(params); got != "StoreUserRequest" {
		t.Fatalf("got %q", got)
	}
	if got := findFormRequestParam([]methodParam{{name: "id", typeName: "int"}}); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
