//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractPairStringValue_KeyMismatch 테스트
package express

import "testing"

func TestExtractPairStringValue_KeyMismatch(t *testing.T) {
	fi := mustParse(t, []byte(`const o = { foo: 'bar' };`))
	if got := extractPairStringValue(firstObject(t, fi), fi.Src, "path"); got != "" {
		t.Fatalf("got %q", got)
	}
}
