//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestExtractDRFFieldType_Attribute 테스트
package django

import "testing"

func TestExtractDRFFieldType_Attribute(t *testing.T) {
	src := []byte("x = serializers.CharField()\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	call := firstCall(root)
	if call == nil {
		t.Fatal("no call")
	}
	if got := extractDRFFieldType(call, src); got != "CharField" {
		t.Fatalf("got %q, want CharField", got)
	}
}
