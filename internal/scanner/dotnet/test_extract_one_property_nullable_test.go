//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestExtractOneProperty_Nullable 테스트
package dotnet

import "testing"

func TestExtractOneProperty_Nullable(t *testing.T) {
	root, src := parseCS(t, `class C { public int? Age { get; set; } }`)
	props := findAllByType(root, "property_declaration")
	f := extractOneProperty(props[0], src)
	if f.Name != "Age" || !f.Nullable {
		t.Fatalf("got %+v", f)
	}
}
