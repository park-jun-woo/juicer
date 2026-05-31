//ff:func feature=scan type=test topic=dotnet control=sequence
//ff:what resultArgType argument_list 첫 인자의 타입 추출 테스트
package dotnet

import "testing"

func TestResultArgType(t *testing.T) {
	root, err := parseCSharp([]byte(`class C { void M() { return Ok(new UserDto()); } }`))
	if err != nil {
		t.Fatal(err)
	}
	src := []byte(`class C { void M() { return Ok(new UserDto()); } }`)
	args := findAllByType(root, "argument_list")
	if len(args) == 0 {
		t.Fatal("no argument_list")
	}
	name, isArr := resultArgType(args[0], src)
	if name != "UserDto" || isArr {
		t.Errorf("got (%q,%v)", name, isArr)
	}
	// nil args
	if _, ok := resultArgType(nil, src); ok {
		t.Error("nil args should be false")
	}
}
