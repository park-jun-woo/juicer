//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestApplySerdeAttrs_Round5 테스트
package actix

import "testing"

func TestApplySerdeAttrs_Round5(t *testing.T) {
	name, nullable := applySerdeAttrs([]serdeAttr{{rename: "user_name"}, {hasDefault: true}}, "userName", false)
	if name != "user_name" || !nullable {
		t.Fatalf("got %q %v", name, nullable)
	}
	name2, nullable2 := applySerdeAttrs(nil, "x", false)
	if name2 != "x" || nullable2 {
		t.Fatalf("no attrs: %q %v", name2, nullable2)
	}
}
