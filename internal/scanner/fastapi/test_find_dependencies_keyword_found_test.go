//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestFindDependenciesKeyword_Found 테스트
package fastapi

import "testing"

func TestFindDependenciesKeyword_Found(t *testing.T) {
	args, src := firstArgList(t, []byte("f('/x', dependencies=[Depends(a), Depends(b)])\n"))
	deps := findDependenciesKeyword(args, src)
	if len(deps) != 2 || deps[0] != "a" || deps[1] != "b" {
		t.Fatalf("got %v", deps)
	}
}
