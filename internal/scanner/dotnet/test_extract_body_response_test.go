//ff:func feature=scan type=test topic=dotnet control=sequence
//ff:what extractBodyResponse 메서드 block에서 응답 추출 및 무body 처리 테스트
package dotnet

import "testing"

func TestExtractBodyResponse(t *testing.T) {
	src := []byte(`class C { IActionResult M() { return Ok(new Dto()); } }`)
	root, _ := parseCSharp(src)
	method := findAllByType(root, "method_declaration")[0]
	res := extractBodyResponse(method, src)
	if !res.found || res.status != "200" || res.typeName != "Dto" {
		t.Errorf("got %+v", res)
	}

	// abstract method (no block) -> not found
	src2 := []byte(`abstract class C { abstract IActionResult M(); }`)
	root2, _ := parseCSharp(src2)
	m2 := findAllByType(root2, "method_declaration")[0]
	if extractBodyResponse(m2, src2).found {
		t.Error("no block should be not found")
	}
}
