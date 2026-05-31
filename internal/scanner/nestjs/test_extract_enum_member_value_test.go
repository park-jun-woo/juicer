//ff:func feature=scan type=test topic=nestjs control=sequence
//ff:what extractEnumMemberValue 이름별 enum 멤버 값 해석(값/무값/부재) 테스트
package nestjs

import "testing"

func TestExtractEnumMemberValue(t *testing.T) {
	src := []byte(`enum RouteKey { Asset = 'assets', Bare }`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	if v, ok := extractEnumMemberValue(root, src, "RouteKey", "Asset"); !ok || v != "assets" {
		t.Errorf("Asset: (%q,%v)", v, ok)
	}
	// valueless member -> key name
	if v, ok := extractEnumMemberValue(root, src, "RouteKey", "Bare"); !ok || v != "Bare" {
		t.Errorf("Bare: (%q,%v)", v, ok)
	}
	// missing member
	if _, ok := extractEnumMemberValue(root, src, "RouteKey", "None"); ok {
		t.Error("missing member should be false")
	}
	// missing enum
	if _, ok := extractEnumMemberValue(root, src, "Other", "Asset"); ok {
		t.Error("missing enum should be false")
	}
}
