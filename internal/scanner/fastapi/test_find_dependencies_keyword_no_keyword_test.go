//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestFindDependenciesKeyword_NoKeyword 테스트
package fastapi

import "testing"

func TestFindDependenciesKeyword_NoKeyword(t *testing.T) {
	args, src := firstArgList(t, []byte("f('/x')\n"))
	if deps := findDependenciesKeyword(args, src); deps != nil {
		t.Fatalf("got %v", deps)
	}
}
