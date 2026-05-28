//ff:func feature=ddl type=test control=iteration dimension=1
//ff:what stripInlineComments가 /* ... */ 블록 주석을 제거하는지 테스트
package ddl

import "testing"

func TestStripInlineComments_Block(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "inline block comment",
			in:   "id INT /* primary key */,\nname TEXT",
			want: "id INT ,\nname TEXT",
		},
		{
			name: "multi-line block comment in body",
			in:   "id INT,\n/* this column\n   is special */\nname TEXT",
			want: "id INT,\n\nname TEXT",
		},
		{
			name: "block and line comments together",
			in:   "id INT /* pk */, -- also pk\nname TEXT",
			want: "id INT , \nname TEXT",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := stripInlineComments(tt.in)
			if got != tt.want {
				t.Fatalf("got %q, want %q", got, tt.want)
			}
		})
	}
}
