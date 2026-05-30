//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestDetectGlobalPrefixInFile_Missing 테스트
package nestjs

import "testing"

func TestDetectGlobalPrefixInFile_Missing(t *testing.T) {
	_, found := detectGlobalPrefixInFile("/no/such.ts")
	if found {
		t.Fatal("expected not found for missing file")
	}
}
