//ff:func feature=ddl type=test control=iteration dimension=1
//ff:what stripLeadingComments가 /* ... */ 블록 주석을 제거하는지 테스트
package ddl

import "testing"

func TestStripLeadingComments_Block(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "multi-line block comment",
			in:   "/* migration\n   v1 */\nCREATE TABLE t (id INT)",
			want: "CREATE TABLE t (id INT)",
		},
		{
			name: "inline block comment",
			in:   "/* comment */ CREATE TABLE t (id INT)",
			want: "CREATE TABLE t (id INT)",
		},
		{
			name: "block then line comment",
			in:   "/* block */\n-- line\nCREATE TABLE t (id INT)",
			want: "CREATE TABLE t (id INT)",
		},
		{
			name: "empty lines then block",
			in:   "\n\n/* block */\nCREATE TABLE t (id INT)",
			want: "CREATE TABLE t (id INT)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := stripLeadingComments(tt.in)
			if got != tt.want {
				t.Fatalf("got %q, want %q", got, tt.want)
			}
		})
	}
}
