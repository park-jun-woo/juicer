//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestExtractKeywordArg 테스트
package flask

import "testing"

func TestExtractKeywordArg(t *testing.T) {
	args, src := argListOf(t, `f("pos", url_prefix="/api", count=5)`+"\n")
	if got := extractKeywordArg(args, "url_prefix", src); got != "/api" {
		t.Fatalf("url_prefix = %q", got)
	}

	if got := extractKeywordArg(args, "count", src); got != "" {
		t.Fatalf("count should be empty (non-string), got %q", got)
	}

	if got := extractKeywordArg(args, "missing", src); got != "" {
		t.Fatalf("missing should be empty, got %q", got)
	}
}
