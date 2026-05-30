//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractSpecImportName_None 테스트
package express

import "testing"

func TestExtractSpecImportName_None(t *testing.T) {

	fi := mustParse(t, []byte(`f("s");`))
	args := findChildByType(firstCallExpr(t, fi), "arguments")
	if args == nil {
		t.Fatal("no arguments")
	}
	if got := extractSpecImportName(args, fi.Src); got != "" {
		t.Fatalf("got %q", got)
	}
}
