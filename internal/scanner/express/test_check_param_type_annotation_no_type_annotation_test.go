//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestCheckParamTypeAnnotation_NoTypeAnnotation 테스트
package express

import "testing"

func TestCheckParamTypeAnnotation_NoTypeAnnotation(t *testing.T) {

	fi := mustParse(t, []byte("function f(r) {}\n"))
	param := firstParamOfType(fi.Root, "required_parameter")
	if param == nil {
		t.Fatal("no required_parameter")
	}
	routers := map[string]bool{}
	checkParamTypeAnnotation(param, fi.Src, routers)
	if len(routers) != 0 {
		t.Fatalf("expected no routers for untyped param, got %v", routers)
	}
}
