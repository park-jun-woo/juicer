//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractStatusFromChain_EmptyArgs 테스트
package express

import "testing"

func TestExtractStatusFromChain_EmptyArgs(t *testing.T) {
	fi := mustParse(t, []byte(`res.status().json({});`))
	if got := extractStatusFromChain(outermostCall(fi), fi.Src); got != "" {
		t.Fatalf("got %q", got)
	}
}
