//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestProcessServiceCallArgs_NoCall 테스트
package actix

import "testing"

func TestProcessServiceCallArgs_NoCall(t *testing.T) {

	src := []byte(`fn f() { cfg.service(handler); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	call := findCallByFuncSuffix(root, src, ".service")
	args := findChildByType(call, "arguments")
	var routes []builderRoute
	processServiceCallArgs(args, src, "", &routes)
	if len(routes) != 0 {
		t.Fatalf("expected no routes, got %+v", routes)
	}
}
