//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestProcessServiceArg_Other 테스트
package actix

import "testing"

func TestProcessServiceArg_Other(t *testing.T) {

	src := []byte(`fn f() { other::thing("/x"); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	call := findCallByFuncSuffix(root, src, "::thing")
	if call == nil {
		t.Fatal("no call")
	}
	var routes []builderRoute
	processServiceArg(call, src, "/api", &routes)
	if len(routes) != 0 {
		t.Fatalf("expected no routes, got %+v", routes)
	}
}
