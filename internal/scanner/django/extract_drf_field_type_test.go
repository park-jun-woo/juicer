//ff:func feature=scan type=test control=sequence topic=django
//ff:what extractDRFFieldType — serializer 필드 타입명 추출 분기를 검증
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

func TestExtractDRFFieldType_Neither(t *testing.T) {
	// Module root has neither attribute nor identifier direct child.
	src := []byte("x = 1\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	if got := extractDRFFieldType(root, src); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
