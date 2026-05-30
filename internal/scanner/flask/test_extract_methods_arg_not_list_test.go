//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestExtractMethodsArg_NotList 테스트
package flask

import "testing"

func TestExtractMethodsArg_NotList(t *testing.T) {

	args, src := argListOf(t, `route('/x', methods=ALLOWED)`+"\n")
	if got := extractMethodsArg(args, src); got != nil {
		t.Fatalf("expected nil for non-list methods, got %v", got)
	}
}
