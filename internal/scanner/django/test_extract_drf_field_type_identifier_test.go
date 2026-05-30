//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestExtractDRFFieldType_Identifier 테스트
package django

import "testing"

func TestExtractDRFFieldType_Identifier(t *testing.T) {
	src := []byte("x = CharField()\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	call := firstCall(root)
	if got := extractDRFFieldType(call, src); got != "CharField" {
		t.Fatalf("got %q, want CharField", got)
	}
}
