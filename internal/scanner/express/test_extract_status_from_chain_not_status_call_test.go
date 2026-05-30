//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractStatusFromChain_NotStatusCall 테스트
package express

import "testing"

func TestExtractStatusFromChain_NotStatusCall(t *testing.T) {

	fi := mustParse(t, []byte(`foo().json({});`))
	if got := extractStatusFromChain(outermostCall(fi), fi.Src); got != "" {
		t.Fatalf("got %q", got)
	}
}
