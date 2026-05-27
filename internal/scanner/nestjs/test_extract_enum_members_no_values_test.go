//ff:func feature=scan type=test control=iteration dimension=1 topic=nestjs
//ff:what extractEnumMembers 값 없는 enum 테스트
package nestjs

import "testing"

func TestExtractEnumMembers_NoValues(t *testing.T) {
	src := []byte(`enum Direction { Up, Down, Left, Right }`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	vals := extractEnumMembers(root, src, "Direction")
	if len(vals) != 4 {
		t.Fatalf("expected 4 members, got %d: %v", len(vals), vals)
	}
	tests := []struct {
		index int
		want  string
	}{
		{0, "Up"},
		{1, "Down"},
		{2, "Left"},
		{3, "Right"},
	}
	for _, tt := range tests {
		if vals[tt.index] != tt.want {
			t.Errorf("vals[%d] = %q, want %q", tt.index, vals[tt.index], tt.want)
		}
	}
}
