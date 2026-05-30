//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestMatchHTTPAttribute_Round5 테스트
package dotnet

import "testing"

func TestMatchHTTPAttribute_Round5(t *testing.T) {
	root, src := parseCS(t, `class C { [HttpPost("/items")] void M() {} }`)
	attrList := firstOfType(t, root, "attribute_list")
	method, path, ok := matchHTTPAttribute(attrList, src)
	if !ok || method != "POST" || path != "/items" {
		t.Fatalf("got %q %q %v", method, path, ok)
	}

	root2, src2 := parseCS(t, `class C { [Serializable] void M() {} }`)
	attrList2 := firstOfType(t, root2, "attribute_list")
	if _, _, ok := matchHTTPAttribute(attrList2, src2); ok {
		t.Fatal("expected no http attribute")
	}
}
