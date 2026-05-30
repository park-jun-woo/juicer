//ff:func feature=scan type=test control=sequence topic=flask
//ff:what extractKeywordArg 테스트
package flask

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func argListOf(t *testing.T, src string) (*sitter.Node, []byte) {
	t.Helper()
	b := []byte(src)
	root, err := parsePython(b)
	if err != nil {
		t.Fatal(err)
	}
	args := findAllByType(root, "argument_list")
	if len(args) == 0 {
		t.Fatal("no argument_list")
	}
	return args[0], b
}

func TestExtractKeywordArg(t *testing.T) {
	args, src := argListOf(t, `f("pos", url_prefix="/api", count=5)`+"\n")
	if got := extractKeywordArg(args, "url_prefix", src); got != "/api" {
		t.Fatalf("url_prefix = %q", got)
	}
	// keyword present but non-string value -> ""
	if got := extractKeywordArg(args, "count", src); got != "" {
		t.Fatalf("count should be empty (non-string), got %q", got)
	}
	// missing keyword -> ""
	if got := extractKeywordArg(args, "missing", src); got != "" {
		t.Fatalf("missing should be empty, got %q", got)
	}
}
