//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what findDependenciesKeyword: dependencies 리스트 추출 / 키 불일치 / 없음
package fastapi

import "testing"

func TestFindDependenciesKeyword_Found(t *testing.T) {
	args, src := firstArgList(t, []byte("f('/x', dependencies=[Depends(a), Depends(b)])\n"))
	deps := findDependenciesKeyword(args, src)
	if len(deps) != 2 || deps[0] != "a" || deps[1] != "b" {
		t.Fatalf("got %v", deps)
	}
}

func TestFindDependenciesKeyword_OtherKeyword(t *testing.T) {
	// keyword args present but none named "dependencies"
	args, src := firstArgList(t, []byte("f('/x', status_code=200)\n"))
	if deps := findDependenciesKeyword(args, src); deps != nil {
		t.Fatalf("got %v", deps)
	}
}

func TestFindDependenciesKeyword_NonListValue(t *testing.T) {
	// dependencies= with a non-list value -> listNode nil -> skipped
	args, src := firstArgList(t, []byte("f('/x', dependencies=common_deps)\n"))
	if deps := findDependenciesKeyword(args, src); deps != nil {
		t.Fatalf("got %v", deps)
	}
}

func TestFindDependenciesKeyword_NoKeyword(t *testing.T) {
	args, src := firstArgList(t, []byte("f('/x')\n"))
	if deps := findDependenciesKeyword(args, src); deps != nil {
		t.Fatalf("got %v", deps)
	}
}
