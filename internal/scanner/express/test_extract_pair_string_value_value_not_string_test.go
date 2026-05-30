//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractPairStringValue_ValueNotString 테스트
package express

import "testing"

func TestExtractPairStringValue_ValueNotString(t *testing.T) {
	fi := mustParse(t, []byte(`const o = { path: somevar };`))
	if got := extractPairStringValue(firstObject(t, fi), fi.Src, "path"); got != "" {
		t.Fatalf("got %q", got)
	}
}
