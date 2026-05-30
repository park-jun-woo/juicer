//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractPairIdentValue_Found 테스트
package express

import "testing"

func TestExtractPairIdentValue_Found(t *testing.T) {
	fi := mustParse(t, []byte(`const o = { route: userRoute, other: x };`))
	if got := extractPairIdentValue(firstObject(t, fi), fi.Src, "route"); got != "userRoute" {
		t.Fatalf("got %q", got)
	}
}
