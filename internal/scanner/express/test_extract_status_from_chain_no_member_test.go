//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractStatusFromChain_NoMember 테스트
package express

import "testing"

func TestExtractStatusFromChain_NoMember(t *testing.T) {
	fi := mustParse(t, []byte(`foo();`))
	if got := extractStatusFromChain(firstCallExpr(t, fi), fi.Src); got != "" {
		t.Fatalf("got %q", got)
	}
}
