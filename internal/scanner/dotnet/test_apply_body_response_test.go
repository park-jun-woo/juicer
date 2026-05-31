//ff:func feature=scan type=test topic=dotnet control=sequence
//ff:what applyBodyResponse 추출 응답을 endpointInfo에 반영(기존 status 보존) 테스트
package dotnet

import "testing"

func TestApplyBodyResponse(t *testing.T) {
	src := []byte(`class C { IActionResult M() { return Ok(new Dto()); } }`)
	root, _ := parseCSharp(src)
	method := findAllByType(root, "method_declaration")[0]

	ep := &endpointInfo{}
	applyBodyResponse(method, src, ep)
	if ep.statusCode != "200" || ep.returnType != "Dto" {
		t.Errorf("ep: %+v", ep)
	}

	// existing statusCode preserved
	ep2 := &endpointInfo{statusCode: "201"}
	applyBodyResponse(method, src, ep2)
	if ep2.statusCode != "201" {
		t.Errorf("existing status must be preserved: %q", ep2.statusCode)
	}
	if ep2.returnType != "Dto" {
		t.Errorf("returnType should still set: %q", ep2.returnType)
	}

	// no response -> no-op
	src2 := []byte(`class C { void M() { var x = 1; } }`)
	root2, _ := parseCSharp(src2)
	m2 := findAllByType(root2, "method_declaration")[0]
	ep3 := &endpointInfo{}
	applyBodyResponse(m2, src2, ep3)
	if ep3.statusCode != "" {
		t.Errorf("no response should be no-op: %+v", ep3)
	}
}
