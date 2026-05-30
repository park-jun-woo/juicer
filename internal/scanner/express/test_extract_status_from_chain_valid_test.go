//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractStatusFromChain_Valid 테스트
package express

import "testing"

func TestExtractStatusFromChain_Valid(t *testing.T) {
	fi := mustParse(t, []byte(`res.status(201).json({});`))
	if got := extractStatusFromChain(outermostCall(fi), fi.Src); got != "201" {
		t.Fatalf("got %q", got)
	}
}
