//ff:func feature=scan type=test control=iteration dimension=1 topic=nestjs
//ff:what extractEnumMembers 문자열 값 enum 테스트
package nestjs

import "testing"

func TestExtractEnumMembers_StringValues(t *testing.T) {
	src := []byte(`enum TaskStatus { OPEN = 'open', IN_PROGRESS = 'in_progress', DONE = 'done' }`)
	root, err := parseTypeScript(src)
	if err != nil {
		t.Fatal(err)
	}
	vals := extractEnumMembers(root, src, "TaskStatus")
	if len(vals) != 3 {
		t.Fatalf("expected 3 members, got %d: %v", len(vals), vals)
	}
	tests := []struct {
		index int
		want  string
	}{
		{0, "open"},
		{1, "in_progress"},
		{2, "done"},
	}
	for _, tt := range tests {
		if vals[tt.index] != tt.want {
			t.Errorf("vals[%d] = %q, want %q", tt.index, vals[tt.index], tt.want)
		}
	}
}
