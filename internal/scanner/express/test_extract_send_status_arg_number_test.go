//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractSendStatusArg_Number 테스트
package express

import "testing"

func TestExtractSendStatusArg_Number(t *testing.T) {
	fi := mustParse(t, []byte(`res.sendStatus(204);`))
	if got := extractSendStatusArg(firstCallExpr(t, fi), fi.Src); got != "204" {
		t.Fatalf("got %q", got)
	}
}
