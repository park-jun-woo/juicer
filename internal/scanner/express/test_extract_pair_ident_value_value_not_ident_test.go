//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractPairIdentValue_ValueNotIdent 테스트
package express

import "testing"

func TestExtractPairIdentValue_ValueNotIdent(t *testing.T) {

	fi := mustParse(t, []byte(`const o = { route: '/literal' };`))
	if got := extractPairIdentValue(firstObject(t, fi), fi.Src, "route"); got != "" {
		t.Fatalf("got %q", got)
	}
}
