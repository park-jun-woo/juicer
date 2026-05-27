//ff:func feature=scan type=test control=iteration dimension=1 topic=nestjs
//ff:what extractEnumMembers 숫자 값 enum 테스트
package nestjs

import "testing"

func TestExtractEnumMembers_NumericValues(t *testing.T) {
	src := []byte(`enum Priority { LOW = 1, MED = 2, HIGH = 3 }`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	vals := extractEnumMembers(root, src, "Priority")
	if len(vals) != 3 {
		t.Fatalf("expected 3 members, got %d: %v", len(vals), vals)
	}
	// numeric enum_assignment has no string child, so falls back to property_identifier key name
	tests := []struct {
		index int
		want  string
	}{
		{0, "LOW"},
		{1, "MED"},
		{2, "HIGH"},
	}
	for _, tt := range tests {
		if vals[tt.index] != tt.want {
			t.Errorf("vals[%d] = %q, want %q", tt.index, vals[tt.index], tt.want)
		}
	}
}
