//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestFindDependenciesKeyword_OtherKeyword 테스트
package fastapi

import "testing"

func TestFindDependenciesKeyword_OtherKeyword(t *testing.T) {

	args, src := firstArgList(t, []byte("f('/x', status_code=200)\n"))
	if deps := findDependenciesKeyword(args, src); deps != nil {
		t.Fatalf("got %v", deps)
	}
}
