//ff:func feature=scan type=test topic=dotnet control=sequence
//ff:what matchStatusCodeInvocation StatusCode(404, body) 정수 status + 본문 타입 추출 테스트
package dotnet

import "testing"

func TestMatchStatusCodeInvocation(t *testing.T) {
	src := []byte(`class C { void M() { return StatusCode(404, new ErrorDto()); } }`)
	root, err := parseCSharp(src)
	if err != nil {
		t.Fatal(err)
	}
	args := findAllByType(root, "argument_list")[0]
	res := matchStatusCodeInvocation(args, src)
	if !res.found || res.status != "404" || res.typeName != "ErrorDto" {
		t.Errorf("got %+v", res)
	}
	// nil args
	if matchStatusCodeInvocation(nil, src).found {
		t.Error("nil should be not found")
	}
	// non-integer first arg
	src2 := []byte(`class C { void M() { return StatusCode(code); } }`)
	root2, _ := parseCSharp(src2)
	a2 := findAllByType(root2, "argument_list")[0]
	if matchStatusCodeInvocation(a2, src2).found {
		t.Error("non-integer status should be not found")
	}
}
