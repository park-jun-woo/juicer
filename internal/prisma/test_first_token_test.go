//ff:func feature=prisma type=test control=iteration dimension=1 topic=prisma
//ff:what firstToken 선두 공백 토큰/나머지 분리 테스트
package prisma

import "testing"

func TestFirstToken(t *testing.T) {
	cases := []struct {
		in       string
		wantTok  string
		wantRest string
	}{
		{"  abc def", "abc", " def"},
		{"abc", "abc", ""},
		{"\t\tx y", "x", " y"},
		{"   ", "", ""},
		{"", "", ""},
	}
	for _, c := range cases {
		tok, rest := firstToken(c.in)
		if tok != c.wantTok || rest != c.wantRest {
			t.Errorf("firstToken(%q) = (%q,%q), want (%q,%q)", c.in, tok, rest, c.wantTok, c.wantRest)
		}
	}
}
