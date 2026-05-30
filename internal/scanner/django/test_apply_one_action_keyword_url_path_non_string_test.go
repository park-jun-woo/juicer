//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestApplyOneActionKeyword_UrlPathNonString 테스트
package django

import "testing"

func TestApplyOneActionKeyword_UrlPathNonString(t *testing.T) {

	src := []byte(`x = action(url_path=some_var)` + "\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	kw := keywordArgs(root)
	if len(kw) == 0 {
		t.Fatal("no keyword_argument")
	}
	ai := &actionInfo{}
	applyOneActionKeyword(ai, "url_path", kw[0], src)
	if ai.urlPath != "" {
		t.Errorf("expected empty urlPath for non-string value, got %q", ai.urlPath)
	}
}
