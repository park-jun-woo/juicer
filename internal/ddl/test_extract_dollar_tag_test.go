//ff:func feature=ddl type=test control=iteration dimension=1
//ff:what extractDollarTag 단위 테스트
package ddl

import "testing"

func TestExtractDollarTag(t *testing.T) {
	tests := []struct {
		name string
		in   string
		pos  int
		want string
	}{
		{name: "empty tag $$", in: "$$body$$", pos: 0, want: "$$"},
		{name: "named tag", in: "$fn$body$fn$", pos: 0, want: "$fn$"},
		{name: "underscore tag", in: "$_tag$body$_tag$", pos: 0, want: "$_tag$"},
		{name: "not a tag - digit start", in: "$1x$", pos: 0, want: ""},
		{name: "not a tag - no closing $", in: "$abc", pos: 0, want: ""},
		{name: "single $", in: "$", pos: 0, want: ""},
		{name: "dollar in middle", in: "foo $$bar$$", pos: 4, want: "$$"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runes := []rune(tt.in)
			got := extractDollarTag(runes, tt.pos)
			if got != tt.want {
				t.Fatalf("got %q, want %q", got, tt.want)
			}
		})
	}
}
