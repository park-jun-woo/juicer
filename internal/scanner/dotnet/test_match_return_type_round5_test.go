//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestMatchReturnType_Round5 테스트
package dotnet

import "testing"

func TestMatchReturnType_Round5(t *testing.T) {

	root, src := parseCS(t, `class C { ActionResult<UserDto> M() { return Ok(); } }`)
	var ep endpointInfo
	gens := findAllByType(root, "generic_name")
	if len(gens) == 0 {
		t.Fatal("no generic_name")
	}
	if !matchReturnType(gens[0], src, &ep) {
		t.Fatal("expected match for generic_name")
	}
	if ep.returnType != "UserDto" {
		t.Fatalf("returnType: %q", ep.returnType)
	}

	root2, src2 := parseCS(t, `class C { void M() {} }`)
	pre := findAllByType(root2, "predefined_type")
	var ep2 endpointInfo
	if !matchReturnType(pre[0], src2, &ep2) {
		t.Fatal("void should match")
	}
}
