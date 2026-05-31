//ff:func feature=scan type=test topic=dotnet control=sequence
//ff:what matchReturnStatement return 문의 invocation에서 bodyResponse 추출 테스트
package dotnet

import "testing"

func TestMatchReturnStatement(t *testing.T) {
	src := []byte(`class C { void M() { return Ok(new Dto()); } }`)
	root, _ := parseCSharp(src)
	ret := findAllByType(root, "return_statement")[0]
	res := matchReturnStatement(ret, src)
	if !res.found || res.status != "200" {
		t.Errorf("got %+v", res)
	}

	// return without invocation
	src2 := []byte(`class C { void M() { return x; } }`)
	root2, _ := parseCSharp(src2)
	ret2 := findAllByType(root2, "return_statement")[0]
	if matchReturnStatement(ret2, src2).found {
		t.Error("non-invocation return should be not found")
	}
}
