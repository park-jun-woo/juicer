//ff:func feature=scan type=test control=sequence topic=supafunc
//ff:what TestCollectMethodsFromCondition 테스트
package supafunc

import "testing"

func TestCollectMethodsFromCondition(t *testing.T) {
	fi := mustParse(t, []byte(`if (req.method === "GET" || req.method === "POST") {}`))
	conds := findAllByType(fi.Root, "parenthesized_expression")
	if len(conds) == 0 {
		t.Fatal("no condition")
	}
	methods := collectMethodsFromCondition(conds[0], fi.Src)
	if len(methods) != 2 {
		t.Fatalf("got %v", methods)
	}
}
