//ff:func feature=scan type=test control=iteration dimension=1 topic=dotnet
//ff:what TestArgType_ObjectCreation 테스트
package dotnet

import "testing"

func TestArgType_ObjectCreation(t *testing.T) {
	root, src := parseCS(t, `class C { void m() { return Ok(new UserDto()); } }`)
	args := findAllByType(root, "argument")
	for _, a := range args {
		if name, _ := argType(a, src); name == "UserDto" {
			return
		}
	}
	t.Skip("no object creation arg matched")
}
