//ff:func feature=scan type=test control=sequence
//ff:what TestExtractPathString_Other 테스트
package fiber

import "testing"

func TestExtractPathString_Other(t *testing.T) {

	_, ok := extractPathFor(t, "pathVar")
	if ok {
		t.Fatal("ident should be false")
	}
}
