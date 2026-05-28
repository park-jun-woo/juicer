//ff:func feature=ddl type=test control=iteration dimension=1
//ff:what stripBlockComments 단위 테스트
package ddl

import "testing"

func TestStripBlockComments(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "no comments",
			in:   "CREATE TABLE t (id INT)",
			want: "CREATE TABLE t (id INT)",
		},
		{
			name: "single inline",
			in:   "id INT /* pk */, name TEXT",
			want: "id INT , name TEXT",
		},
		{
			name: "multi-line",
			in:   "/* line1\nline2 */\nCREATE TABLE t (id INT)",
			want: "\nCREATE TABLE t (id INT)",
		},
		{
			name: "multiple block comments",
			in:   "/* a */ x /* b */ y",
			want: " x  y",
		},
		{
			name: "unclosed block comment",
			in:   "/* unclosed",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := stripBlockComments(tt.in)
			if got != tt.want {
				t.Fatalf("got %q, want %q", got, tt.want)
			}
		})
	}
}
