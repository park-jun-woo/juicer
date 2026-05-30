//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractSendStatusArg_Empty 테스트
package express

import "testing"

func TestExtractSendStatusArg_Empty(t *testing.T) {
	fi := mustParse(t, []byte(`res.sendStatus();`))
	if got := extractSendStatusArg(firstCallExpr(t, fi), fi.Src); got != "" {
		t.Fatalf("got %q", got)
	}
}
