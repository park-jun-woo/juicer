//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractZodValidators_MixedArgs 테스트
package hono

import "testing"

func TestExtractZodValidators_MixedArgs(t *testing.T) {

	fi := mustParse(t, []byte(`app.post("/x", zValidator("json", s), other(), handler);`+"\n"))
	args := findAllByType(fi.Root, "arguments")[0]
	nodes := collectArgNodes(args)
	vs := extractZodValidators(nodes, fi.Src)
	if len(vs) != 1 || vs[0].Target != "json" || vs[0].SchemaName != "s" {
		t.Fatalf("got %+v", vs)
	}
}
