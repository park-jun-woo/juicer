//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestExtractSerializersFromFile_None 테스트
package django

import "testing"

func TestExtractSerializersFromFile_None(t *testing.T) {
	fi := newTestFileInfo(t, "x = 1\n")
	if s := extractSerializersFromFile(fi); len(s) != 0 {
		t.Fatalf("expected none, got %d", len(s))
	}
}
