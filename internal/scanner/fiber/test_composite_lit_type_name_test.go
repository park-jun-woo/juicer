//ff:func feature=scan type=test control=iteration dimension=1
//ff:what TestCompositeLitTypeName — composite-literal body 문자열 파싱 테스트
package fiber

import "testing"

func TestCompositeLitTypeName(t *testing.T) {
	cases := []struct {
		body, wantDisplay, wantBase string
	}{
		{"Book{}", "Book", "Book"},
		{"[]Book{}", "[]Book", "Book"},
		{"pkg.Book{}", "", ""},
		{"books", "", ""},
		{"map[string]any{}", "", ""},
	}
	for _, c := range cases {
		display, base := compositeLitTypeName(c.body)
		if display != c.wantDisplay || base != c.wantBase {
			t.Errorf("compositeLitTypeName(%q) = (%q,%q), want (%q,%q)", c.body, display, base, c.wantDisplay, c.wantBase)
		}
	}
}
