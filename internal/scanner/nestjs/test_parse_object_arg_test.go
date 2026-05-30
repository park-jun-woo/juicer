//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestParseObjectArg 테스트
package nestjs

import "testing"

func TestParseObjectArg(t *testing.T) {
	src := []byte(`const o = { path: 'auth', version: '1' };`)
	root, _ := parseTypeScript(src)
	obj := findAllByType(root, "object")[0]
	d := &decoratorInfo{objectProps: map[string]string{}}
	parseObjectArg(obj, src, d)
	if d.arg != "auth" {
		t.Fatalf("arg: %q", d.arg)
	}
	if d.objectProps["version"] != "1" {
		t.Fatalf("props: %v", d.objectProps)
	}
}
