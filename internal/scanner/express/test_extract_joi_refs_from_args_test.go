//ff:func feature=scan type=test topic=express control=sequence
//ff:what extractJoiRefsFromArgs 인자 목록에서 start 이후 validate 참조 수집 테스트
package express

import "testing"

func TestExtractJoiRefsFromArgs(t *testing.T) {
	fi := mustParse(t, []byte(`f(auth(), validate(schemas.body), handler);`))
	outer := findAllByType(fi.Root, "call_expression")[0]
	args := findChildByType(outer, "arguments")
	nodes := collectArgNodes(args)

	// start 0: collects validate(schemas.body)
	refs := extractJoiRefsFromArgs(nodes, fi.Src, 0)
	if len(refs) != 1 || refs[0].Member != "body" {
		t.Fatalf("start0: %+v", refs)
	}
	// start beyond length -> nil
	if r := extractJoiRefsFromArgs(nodes, fi.Src, len(nodes)+1); r != nil {
		t.Errorf("oob start: %+v", r)
	}
}
