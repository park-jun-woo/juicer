//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestExtractKeywordArg_WrongName 테스트
package django

import "testing"

func TestExtractKeywordArg_WrongName(t *testing.T) {
	src := []byte("x = path('a/', view, name='home')\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	args := firstArgumentList(root)
	if got := extractKeywordArg(args, "missing", src); got != "" {
		t.Fatalf("expected empty for missing name, got %q", got)
	}
}
