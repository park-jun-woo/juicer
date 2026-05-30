//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestExtractRecordParam_Round5 테스트
package dotnet

import "testing"

func TestExtractRecordParam_Round5(t *testing.T) {
	root, src := parseCS(t, `public record R(string Name, int? Age);`)
	params := findAllByType(root, "parameter")
	if len(params) < 2 {
		t.Fatalf("expected 2 params, got %d", len(params))
	}
	f0 := extractRecordParam(params[0], src)
	if f0.Name != "Name" || f0.Type != "string" {
		t.Fatalf("param0: %+v", f0)
	}
	f1 := extractRecordParam(params[1], src)
	if f1.Name != "Age" || !f1.Nullable {
		t.Fatalf("param1 should be nullable int: %+v", f1)
	}
}
