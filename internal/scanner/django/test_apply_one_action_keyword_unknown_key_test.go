//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestApplyOneActionKeyword_UnknownKey 테스트
package django

import "testing"

func TestApplyOneActionKeyword_UnknownKey(t *testing.T) {
	src := []byte(`x = action(permission_classes=[A])` + "\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	kw := keywordArgs(root)
	if len(kw) == 0 {
		t.Fatal("no keyword_argument")
	}
	ai := &actionInfo{}
	applyOneActionKeyword(ai, "permission_classes", kw[0], src)
	if ai.detail || len(ai.methods) != 0 || ai.urlPath != "" {
		t.Errorf("unknown key should not modify actionInfo, got %+v", ai)
	}
}
