//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractStatusFromChain_NotNumberArg 테스트
package express

import "testing"

func TestExtractStatusFromChain_NotNumberArg(t *testing.T) {
	fi := mustParse(t, []byte(`res.status(code).json({});`))
	if got := extractStatusFromChain(outermostCall(fi), fi.Src); got != "" {
		t.Fatalf("got %q", got)
	}
}
