//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractZodValidators_None 테스트
package hono

import "testing"

func TestExtractZodValidators_None(t *testing.T) {
	fi := mustParse(t, []byte(`app.get("/x", handler);`+"\n"))
	args := findAllByType(fi.Root, "arguments")[0]
	nodes := collectArgNodes(args)
	if vs := extractZodValidators(nodes, fi.Src); len(vs) != 0 {
		t.Fatalf("expected none, got %+v", vs)
	}
}
