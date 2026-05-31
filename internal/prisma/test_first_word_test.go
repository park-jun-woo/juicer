//ff:func feature=prisma type=test control=iteration dimension=1 topic=prisma
//ff:what firstWord 경계(공백/콤마/괄호/대괄호) 분리 테스트
package prisma

import "testing"

func TestFirstWord(t *testing.T) {
	cases := []struct {
		in       string
		wantWord string
		wantRest string
	}{
		{"abc def", "abc", " def"},
		{"abc,def", "abc", ",def"},
		{"abc)def", "abc", ")def"},
		{"abc]def", "abc", "]def"},
		{"abc\tdef", "abc", "\tdef"},
		{"single", "single", ""},
		{"", "", ""},
	}
	for _, c := range cases {
		w, r := firstWord(c.in)
		if w != c.wantWord || r != c.wantRest {
			t.Errorf("firstWord(%q) = (%q,%q), want (%q,%q)", c.in, w, r, c.wantWord, c.wantRest)
		}
	}
}
