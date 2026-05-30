//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractSpecImportName_IdentifierFallback 테스트
package express

import "testing"

func TestExtractSpecImportName_IdentifierFallback(t *testing.T) {

	fi := mustParse(t, []byte(`f(myArg);`))
	args := findChildByType(firstCallExpr(t, fi), "arguments")
	if args == nil {
		t.Fatal("no arguments")
	}
	if got := extractSpecImportName(args, fi.Src); got != "myArg" {
		t.Fatalf("got %q", got)
	}
}
