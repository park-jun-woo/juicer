//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestParseFile_Missing 테스트
package quarkus

import "testing"

func TestParseFile_Missing(t *testing.T) {
	if _, err := parseFile("/abs", "/no/such.java"); err == nil {
		t.Fatal("expected error")
	}
}
