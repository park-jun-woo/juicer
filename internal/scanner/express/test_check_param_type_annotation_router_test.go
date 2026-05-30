//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestCheckParamTypeAnnotation_Router 테스트
package express

import "testing"

func TestCheckParamTypeAnnotation_Router(t *testing.T) {
	fi := mustParse(t, []byte("function f(r: Router) {}\n"))
	param := firstParamOfType(fi.Root, "required_parameter")
	if param == nil {
		t.Fatal("no required_parameter")
	}
	routers := map[string]bool{}
	checkParamTypeAnnotation(param, fi.Src, routers)
	if !routers["r"] {
		t.Fatalf("expected r registered as router, got %v", routers)
	}
}
