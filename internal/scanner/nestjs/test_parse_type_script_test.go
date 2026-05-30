//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestParseTypeScript 테스트
package nestjs

import "testing"

func TestParseTypeScript(t *testing.T) {
	root, err := parseTypeScript([]byte(`const x = 1;`))
	if err != nil || root == nil {
		t.Fatalf("err: %v", err)
	}
}
