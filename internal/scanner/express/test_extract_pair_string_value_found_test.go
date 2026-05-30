//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractPairStringValue_Found 테스트
package express

import "testing"

func TestExtractPairStringValue_Found(t *testing.T) {
	fi := mustParse(t, []byte(`const o = { path: '/users', other: x };`))
	if got := extractPairStringValue(firstObject(t, fi), fi.Src, "path"); got != "/users" {
		t.Fatalf("got %q", got)
	}
}
