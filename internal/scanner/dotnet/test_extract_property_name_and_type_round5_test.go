//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestExtractPropertyName_And_Type_Round5 테스트
package dotnet

import "testing"

func TestExtractPropertyName_And_Type_Round5(t *testing.T) {
	root, src := parseCS(t, `class C { public string Name { get; set; } }`)
	prop := firstOfType(t, root, "property_declaration")
	if got := extractPropertyName(prop, src); got != "Name" {
		t.Fatalf("name: got %q", got)
	}
	if got := extractPropertyTypeName(prop, src); got != "string" {
		t.Fatalf("type: got %q", got)
	}
}
