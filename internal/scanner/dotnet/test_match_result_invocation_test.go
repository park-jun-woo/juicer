//ff:func feature=scan type=test topic=dotnet control=sequence
//ff:what matchResultInvocation Ok()/StatusCode()/미지원 메서드 분기 테스트
package dotnet

import "testing"

func TestMatchResultInvocation(t *testing.T) {
	// Ok(new Dto()) -> 200
	src := []byte(`class C { void M() { return Ok(new Dto()); } }`)
	root, _ := parseCSharp(src)
	inv := findAllByType(root, "invocation_expression")[0]
	res := matchResultInvocation(inv, src)
	if !res.found || res.status != "200" || res.typeName != "Dto" {
		t.Errorf("Ok: %+v", res)
	}

	// StatusCode branch
	src2 := []byte(`class C { void M() { return StatusCode(403); } }`)
	root2, _ := parseCSharp(src2)
	inv2 := findAllByType(root2, "invocation_expression")[0]
	if res2 := matchResultInvocation(inv2, src2); !res2.found || res2.status != "403" {
		t.Errorf("StatusCode: %+v", res2)
	}

	// unsupported method
	src3 := []byte(`class C { void M() { return Frobnicate(); } }`)
	root3, _ := parseCSharp(src3)
	inv3 := findAllByType(root3, "invocation_expression")[0]
	if matchResultInvocation(inv3, src3).found {
		t.Error("unsupported method should be not found")
	}
}
