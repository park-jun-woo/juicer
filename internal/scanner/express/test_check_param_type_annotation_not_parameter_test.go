//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestCheckParamTypeAnnotation_NotParameter 테스트
package express

import "testing"

func TestCheckParamTypeAnnotation_NotParameter(t *testing.T) {

	fi := mustParse(t, []byte("const x = 1;\n"))
	routers := map[string]bool{}
	checkParamTypeAnnotation(fi.Root, fi.Src, routers)
	if len(routers) != 0 {
		t.Fatalf("expected no routers, got %v", routers)
	}
}
