//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestApplyActionKeywords_NonKeywordArgs 테스트
package django

import "testing"

func TestApplyActionKeywords_NonKeywordArgs(t *testing.T) {

	src := []byte(`x = action('a', 'b')` + "\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	args := firstArgumentList(root)
	ai := &actionInfo{}
	applyActionKeywords(ai, args, src)
	if ai.detail || len(ai.methods) != 0 || ai.urlPath != "" {
		t.Errorf("expected nothing applied, got %+v", ai)
	}
}
