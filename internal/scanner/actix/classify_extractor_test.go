//ff:func feature=scan type=test control=selection topic=actix
//ff:what classifyExtractor — scoped 타입을 extractor 종류로 분류함을 검증
package actix

import "testing"

func TestClassifyExtractor(t *testing.T) {
	cases := map[string]string{
		"web::Path":  "path",
		"web::Json":  "json",
		"web::Query": "query",
		"web::Form":  "form",
		"web::Data":  "", // not an extractor type
	}
	for in, want := range cases {
		if got := classifyExtractor(in); got != want {
			t.Errorf("classifyExtractor(%q) = %q, want %q", in, got, want)
		}
	}
}
