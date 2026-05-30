//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestExtractClassProps 테스트
package dotnet

import "testing"

func TestExtractClassProps(t *testing.T) {
	root, src := parseCS(t, `class UserDto {
		public string Name { get; set; }
		public int Age { get; set; }
	}`)
	cls := findAllByType(root, "class_declaration")[0]
	props := extractClassProps(cls, src)
	if len(props) != 2 || props[0].Name != "Name" {
		t.Fatalf("got %+v", props)
	}
	if props[0].Type != "string" || props[1].Type != "integer" {
		t.Fatalf("types: %+v", props)
	}
}
