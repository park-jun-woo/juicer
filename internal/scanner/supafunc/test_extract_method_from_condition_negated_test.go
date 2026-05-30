//ff:func feature=scan type=test control=sequence topic=supafunc
//ff:what TestExtractMethodFromCondition_Negated 테스트
package supafunc

import "testing"

func TestExtractMethodFromCondition_Negated(t *testing.T) {

	fi := mustParse(t, []byte(`if (req.method !== "POST") {}`))
	conds := findAllByType(fi.Root, "parenthesized_expression")
	if got := extractMethodFromCondition(conds[0], fi.Src); got != "" {
		t.Fatalf("expected empty for negated, got %q", got)
	}
}
