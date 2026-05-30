//ff:func feature=scan type=test control=sequence topic=supafunc
//ff:what TestExtractMethodFromCondition 테스트
package supafunc

import "testing"

func TestExtractMethodFromCondition(t *testing.T) {
	fi := mustParse(t, []byte(`if (req.method === "POST") {}`))
	conds := findAllByType(fi.Root, "parenthesized_expression")
	if len(conds) == 0 {
		t.Fatal("no condition")
	}
	if got := extractMethodFromCondition(conds[0], fi.Src); got != "POST" {
		t.Fatalf("got %q", got)
	}
}
