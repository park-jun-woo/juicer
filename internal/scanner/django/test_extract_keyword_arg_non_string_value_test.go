//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestExtractKeywordArg_NonStringValue 테스트
package django

import "testing"

func TestExtractKeywordArg_NonStringValue(t *testing.T) {
	src := []byte("x = path('a/', view, count=5)\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	args := firstArgumentList(root)
	if got := extractKeywordArg(args, "count", src); got != "" {
		t.Fatalf("expected empty for non-string value, got %q", got)
	}
}
