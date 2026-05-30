//ff:func feature=scan type=test control=sequence topic=express
//ff:what extractSendStatusArg: number인자 / args없음 / 빈args / 비number 분기
package express

import "testing"

func TestExtractSendStatusArg_Number(t *testing.T) {
	fi := mustParse(t, []byte(`res.sendStatus(204);`))
	if got := extractSendStatusArg(firstCallExpr(t, fi), fi.Src); got != "204" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractSendStatusArg_NoArgsNode(t *testing.T) {
	fi := mustParse(t, []byte("res.sendStatus`x`;"))
	if got := extractSendStatusArg(firstCallExpr(t, fi), fi.Src); got != "" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractSendStatusArg_Empty(t *testing.T) {
	fi := mustParse(t, []byte(`res.sendStatus();`))
	if got := extractSendStatusArg(firstCallExpr(t, fi), fi.Src); got != "" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractSendStatusArg_NotNumber(t *testing.T) {
	fi := mustParse(t, []byte(`res.sendStatus(code);`))
	if got := extractSendStatusArg(firstCallExpr(t, fi), fi.Src); got != "" {
		t.Fatalf("got %q", got)
	}
}
