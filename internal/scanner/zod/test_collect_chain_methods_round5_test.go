//ff:func feature=scan type=test control=iteration dimension=1 topic=zod
//ff:what TestCollectChainMethods_Round5 테스트
package zod

import "testing"

func TestCollectChainMethods_Round5(t *testing.T) {
	root, src := parseTS(t, "const s = z.string().min(3).email();")
	call := findFirstCall(t, root)
	methods := CollectChainMethods(call, src)
	if len(methods) == 0 {
		t.Fatalf("expected chain methods, got %+v", methods)
	}

	names := make([]string, len(methods))
	for i, m := range methods {
		names[i] = m.Name
	}
	hasString, hasMin, hasEmail := false, false, false
	for _, n := range names {
		switch n {
		case "string":
			hasString = true
		case "min":
			hasMin = true
		case "email":
			hasEmail = true
		}
	}
	if !hasString || !hasMin || !hasEmail {
		t.Fatalf("missing chain methods: %v", names)
	}
}
