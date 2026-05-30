//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestFindDependenciesKeyword_NonListValue 테스트
package fastapi

import "testing"

func TestFindDependenciesKeyword_NonListValue(t *testing.T) {

	args, src := firstArgList(t, []byte("f('/x', dependencies=common_deps)\n"))
	if deps := findDependenciesKeyword(args, src); deps != nil {
		t.Fatalf("got %v", deps)
	}
}
