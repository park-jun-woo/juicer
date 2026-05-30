//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestFindAttributeInList_Round5 테스트
package dotnet

import "testing"

func TestFindAttributeInList_Round5(t *testing.T) {
	root, src := parseCS(t, `class C { [HttpGet("/x")] public void M() {} }`)
	attrList := firstOfType(t, root, "attribute_list")
	if got := findAttributeInList(attrList, src, "HttpGet"); got == nil {
		t.Fatal("expected HttpGet attribute found")
	}
	if got := findAttributeInList(attrList, src, "Missing"); got != nil {
		t.Fatal("expected nil for missing attribute")
	}
}
